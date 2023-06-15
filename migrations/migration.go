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

	admin := uuid.New()
	buyer := uuid.New()
	seller := uuid.New()

	db.Create(&[]model.Role{{
		ID:             buyer,
		Code:           "RO0001",
		Name:           model.BUYER,
		Scope:          "",
		RolePermission: []*model.RolePermission{},
		OrmModel:       orm.OrmModel{},
	}, {
		ID:             admin,
		Code:           "RO0002",
		Name:           model.ADMIN,
		Scope:          "",
		RolePermission: []*model.RolePermission{},
		OrmModel:       orm.OrmModel{},
	}, {
		ID:             seller,
		Code:           "RO0003",
		Name:           model.SALLER,
		Scope:          "",
		RolePermission: []*model.RolePermission{},
		OrmModel:       orm.OrmModel{},
	}})

	profilePermission := uuid.New()

	db.Create(&[]model.Permission{
		{
			ID:   profilePermission,
			Code: "PM0002",
			Name: "profile",
			Path: "/profile",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.New(),
					RoleID:         buyer,
					PermissionID:   profilePermission,
					PermissionName: "profile",
					PermissionPath: "/profile",
				},
				{
					ID:             uuid.New(),
					RoleID:         seller,
					PermissionID:   profilePermission,
					PermissionName: "profile",
					PermissionPath: "/profile",
				},
				{
					ID:             uuid.New(),
					RoleID:         admin,
					PermissionID:   profilePermission,
					PermissionName: "profile",
					PermissionPath: "/profile",
				},
			},
			OrmModel: orm.OrmModel{},
		},
	})

	// db.Create(&[]model.RolePermission{
	// 	{
	// 		ID:             uuid.New(),
	// 		RoleID:         buyer,
	// 		PermissionID:   profilePermission,
	// 		PermissionName: "profile",
	// 		PermissionPath: "/profile",
	// 		Permission:     model.Permission{},
	// 		OrmModel:       orm.OrmModel{},
	// 	},
	// })
}
