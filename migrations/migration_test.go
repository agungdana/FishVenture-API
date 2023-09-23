package migrations_test

import (
	"log"
	"testing"

	"github.com/e-fish/api/migrations"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/orm"
)

func Test_Migrate(t *testing.T) {
	logger.SetupLogger("true")

	con, err := orm.CreateConnetionDB(migrations.GetConfig().DbConfig)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	// err = migrations.Migrate(con, "initial-data-model")
	err = migrations.Migrate(con, "add-permission")
	if err != nil {
		log.Println("error: ", err)
		return
	}

}
