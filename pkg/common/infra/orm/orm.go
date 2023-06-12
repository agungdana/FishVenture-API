package orm

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

var (
	once  sync.Once
	pools *gormPool
)

type gormPool struct {
	mut sync.Mutex

	pools map[string]*gorm.DB
	txns  map[uuid.UUID]transaction
}

type transaction struct {
	tx *gorm.DB

	//after begin transaction counter += 1
	//commit transaction counter -= 1, commit after counter == 0
	//rollback transaction force transaction
	counter int
}

func getGormPool() *gormPool {
	if pools == nil {
		once.Do(func() {
			pools = &gormPool{
				pools: make(map[string]*gorm.DB),
				txns:  make(map[uuid.UUID]transaction),
			}
		})
	}

	return pools
}

func gormLogger() gormlog.Interface {
	myLogger := logger.GetLogger()
	log := gormlog.New(myLogger.Log, gormlog.Config{
		SlowThreshold:             time.Millisecond,
		Colorful:                  false,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  gormlog.LogLevel(myLogger.Log.Level),
	})

	return log
}

func CreateConnetionDB(conf config.DbConfig) (*gorm.DB, error) {
	var (
		db                 *gorm.DB
		dsn                string
		dsnWithoutPassword string
		exist              bool
		err                error
	)

	switch conf.Driver {
	case "postgres":
		// dsn = "postgres://" + conf.User + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port + "/" + conf.Password
		// dsn = fmt.Sprintf("postgres://%v:%v@%v:%v/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.Password)
		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=UTC", conf.Host, conf.User, conf.Password, conf.Database, conf.Port)
	case "mysql":
		//mysql code dsn here
	default:
		return nil, ErrDriverNotSupported.AttacthDetail(map[string]any{"driveName": conf.Driver, "supportedDrivers": "[postgres]"})
	}

	dsnWithoutPassword = strings.ReplaceAll(dsn, conf.Password, "")

	if db, exist = getGormPool().pools[dsnWithoutPassword]; exist {
		logger.Debug("using an existing connection")
		return db, nil
	}

	logger.Debug("create a new %v connection", conf.Driver)
	switch conf.Driver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLogger(),
		})
	case "mysql":
		//mysql code gorm open here
	default:
		return nil, ErrDriverNotSupported.AttacthDetail(map[string]any{"driveName": conf.Driver, "supportedDrivers": "[postgres]"})
	}

	if err != nil {
		return nil, ErrCreateConnection.AttacthDetail(map[string]any{"error": err})
	}

	if db == nil {
		return nil, ErrCreateConnection.AttacthDetail(map[string]any{"dbConnection": "is empty"})
	}

	getGormPool().pools[dsnWithoutPassword] = db
	logger.Debug("success create connection")
	return db, nil

}

func BeginTxn(ctx context.Context, db *gorm.DB) *gorm.DB {
	txnID, withTransaction := ctxutil.GetTransactionID(ctx)
	if !withTransaction {
		logger.Debug("can")
		return nil
	}

	pool := getGormPool()
	pool.mut.Lock()
	defer pool.mut.Unlock()
	txn, ok := pool.txns[txnID]
	if !ok {
		//create new txn
		txn = transaction{tx: db.Begin()}
	} else {
		//increase counter txn
		txn.counter += 1
	}

	pool.txns[txnID] = txn
	return txn.tx
}

func CommitTxn(ctx context.Context) error {
	txnID, withTransaction := ctxutil.GetTransactionID(ctx)
	if !withTransaction {
		logger.Debug("error empty transaction id")
		return nil
	}

	pool := getGormPool()

	txn, exist := pool.txns[txnID]
	if !exist {
		logger.Debug("error empty transaction")
		return nil
	}

	if txn.counter > 0 {
		txn.counter -= 1
		pool.txns[txnID] = txn
		return nil
	}

	err := txn.tx.Commit().Error
	endTxn(ctx)
	return err
}

func RollbackTxn(ctx context.Context) error {
	txnID, withTransaction := ctxutil.GetTransactionID(ctx)
	if !withTransaction {
		logger.Debug("error empty transaction id")
		return nil
	}

	pool := getGormPool()

	txn, exist := pool.txns[txnID]
	if !exist {
		logger.Debug("error empty transaction")
		return nil
	}

	if txn.counter > 0 {
		txn.counter -= 1
		pool.txns[txnID] = txn
		return nil
	}

	err := txn.tx.Rollback().Error
	endTxn(ctx)
	return err
}

func endTxn(ctx context.Context) error {
	txnID, withTransaction := ctxutil.GetTransactionID(ctx)
	if !withTransaction {
		logger.Debug("error empty transaction id")
		return nil
	}

	pool := getGormPool()
	pool.mut.Lock()
	defer pool.mut.Unlock()

	txn, exist := pool.txns[txnID]
	if !exist {
		logger.Debug("error empty transaction")
		return nil
	}

	if txn.counter > 0 {
		return nil
	}
	delete(pool.txns, txnID)
	return nil
}
