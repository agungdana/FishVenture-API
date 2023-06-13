package migrations

import (
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/auth/model"
	"github.com/google/uuid"
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

	db.Create([]model.Role{{
		ID:             uuid.New(),
		Code:           "RO0001",
		Name:           model.BUYER,
		Scope:          "",
		RolePermission: []*model.RolePermission{},
		OrmModel:       orm.OrmModel{},
	}, {
		ID:             uuid.New(),
		Code:           "RO0002",
		Name:           model.ADMIN,
		Scope:          "",
		RolePermission: []*model.RolePermission{},
		OrmModel:       orm.OrmModel{},
	}, {
		ID:             uuid.New(),
		Code:           "RO0003",
		Name:           model.SALLER,
		Scope:          "",
		RolePermission: []*model.RolePermission{},
		OrmModel:       orm.OrmModel{},
	}})
}
