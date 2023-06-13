package migrations

import (
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/auth/model"
)

func Migrations() {
	db, err := orm.CreateConnetionDB(getConfig().DbConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.RolePermission{},
		&model.UserRole{},
		&model.UserPermission{},
	)
	if err != nil {
		logger.Info("Error Auto Migreate: %v", err)
	}
}
