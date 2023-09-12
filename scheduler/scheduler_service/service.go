package schedulerservice

import (
	"context"
	"sync"
	"time"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/domain/budidaya"
	"github.com/e-fish/api/pkg/domain/pond"
	"github.com/e-fish/api/pkg/domain/transaction"
	"github.com/e-fish/api/pkg/domain/transaction/model"
	"github.com/e-fish/api/pkg/domain/verification"
	schedulerconfig "github.com/e-fish/api/scheduler/scheduler_config"
	"github.com/e-fish/api/scheduler/scheduler_service/internal"
)

type Service struct {
	conf            schedulerconfig.SchedulerConfig
	budidayaRepo    budidaya.Repo
	transactionRepo transaction.Repo
}

func NewService(conf schedulerconfig.SchedulerConfig) Service {
	var (
		service = Service{
			conf: conf,
		}
	)

	verificationRepo, err := verification.NewRepo(conf.BudidayaConfig.DbConfig)
	if err != nil {
		logger.Fatal("failed to create a new repo tenant, can't create scheduler service err: %v", err)
	}

	pondRepo, err := pond.NewRepo(conf.BudidayaConfig.DbConfig, verificationRepo)
	if err != nil {
		logger.Fatal("failed to create a new repo tenant, can't create scheduler service err: %v", err)
	}

	budidayaRepo, err := budidaya.NewRepo(conf.BudidayaConfig.DbConfig, pondRepo)
	if err != nil {
		logger.Fatal("failed to create a new repo tenant, can't create product service err: %v", err)
	}

	transactionRepo, err := transaction.NewRepo(conf.TransactionConfig.DbConfig, budidayaRepo)
	if err != nil {
		logger.Fatal("failed to create a new repo product, can't create product service err: %v", err)
	}

	service.budidayaRepo = budidayaRepo
	service.transactionRepo = transactionRepo

	if conf.CreateProductAfterRun {
		go service.UpdateOrderStatus(1)
	}
	service.Start()

	return service
}

func (s *Service) Start() {
	logger.Debug("Start scheduler")
	go s.cancelOrder()

}

func (s *Service) cancelOrder() {
	shceduler := internal.NewScheduler(s.conf.TimerUpdate)
	shceduler.SetSchedule()

	for {
		<-shceduler.Timer.C
		s.UpdateOrderStatus(1)
		shceduler.SetSchedule()
	}
}

func (s *Service) UpdateOrderStatus(count int) {
	ctx := context.Background()
	ctx = ctxutil.NewRequest(ctx)
	wg := sync.WaitGroup{}

	if count == 10 {
		return
	}

	query := s.transactionRepo.NewQuery()
	orders, err := query.ReadAllOrderActive(ctx)
	if err != nil {
		s.UpdateOrderStatus(count + 1)
		return
	}

	numWorker := make(chan struct{}, 20)
	for _, order := range orders {
		if order == nil {
			continue
		}
		wg.Add(1)
		numWorker <- struct{}{}
		go func(order model.Order) {
			defer func() {
				<-numWorker
				wg.Done()
			}()

			if order.BookingDate == nil {
				return
			}

			lastTime := order.BookingDate.AddDate(0, 0, 2)
			if time.Now().Before(lastTime) {
				return
			}

			command := s.transactionRepo.NewCommand(ctx)
			result, err := command.UpdateCancelOrder(ctx, order.ID)
			if err != nil {
				if err := command.Rollback(ctx); err != nil {
					logger.ErrorWithContext(ctx, "failed rollback transaction update transaction: %v", err)
				}

				return
			}

			if err := command.Commit(ctx); err != nil {
				logger.ErrorWithContext(ctx, "failed commit transaction update transaction: %v", err)
				return
			}

			logger.InfoWithContext(ctx, "Success update order [%v]", result)

		}(*order)
	}

	wg.Wait()

	logger.DebugWithContext(ctx, "########Success update data")
}
