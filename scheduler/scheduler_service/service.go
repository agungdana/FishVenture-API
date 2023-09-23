package schedulerservice

import (
	"context"
	"time"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/domain/budidaya"
	"github.com/e-fish/api/pkg/domain/pond"
	"github.com/e-fish/api/pkg/domain/transaction"
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

	if count == 10 {
		return
	}

	query := s.transactionRepo.NewQuery()
	orders, err := query.ReadAllOrderActive(ctx)
	if err != nil {
		s.UpdateOrderStatus(count + 1)
		return
	}

	for _, order := range orders {
		if order.BookingDate == nil {
			continue
		}
		lastTime := order.BookingDate.AddDate(0, 0, 2)
		if time.Now().Before(lastTime) {
			continue
		}
		command := s.transactionRepo.NewCommand(ctx)
		result, err := command.UpdateCancelOrder(ctx, order.ID)
		if err != nil {
			if err := command.Rollback(ctx); err != nil {
				logger.ErrorWithContext(ctx, "failed rollback transaction update transaction: %v", err)
			}
			continue
		}
		if err := command.Commit(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed commit transaction update transaction: %v", err)
			continue
		}
		logger.InfoWithContext(ctx, "Success update order [%v]", result)
	}
	logger.DebugWithContext(ctx, "########Success update data")
}
