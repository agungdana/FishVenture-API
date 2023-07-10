package migrations

import (
	"fmt"

	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/auth/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Migrations() {
	db, err := orm.CreateConnetionDB(getConfig().DbConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = Migrate(db, "add-permission")
	fmt.Printf("err: %v\n", err)

}

func Migrate(db *gorm.DB, flag string) error {
	switch flag {
	case "initial-data-model":
		err := db.AutoMigrate(
			&User{},
			&Role{},
			&Permission{},
			&RolePermission{},
			&UserRole{},
			&UserPermission{},
			&Team{},
			&Pond{},
			&Berkas{},
			&Pool{},
			&Budidaya{},
			&PriceList{},
			&FishSpecies{},
			&Order{},
			&Country{},
			&Province{},
			&City{},
			&District{},
		)
		if err != nil {
			logger.Info("Error Auto Migreate: %v", err)
		}
	case "add-permission":
		admin := uuid.MustParse("885e314b-b007-4954-8435-f64f7cb02263")
		buyer := uuid.MustParse("013ebbf4-f75c-448a-bded-97b651e8b453")
		seller := uuid.MustParse("0879aaab-0ca1-487b-9c1c-c5805b4403f8")

		db.Save(&[]model.Role{{
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
			Name:           model.SELLER,
			Scope:          "",
			RolePermission: []*model.RolePermission{},
			OrmModel:       orm.OrmModel{},
		}})

		permission := []model.Permission{}

		profilePermission := uuid.MustParse("7f56621d-220e-4628-8c6c-2373ab151862")
		permissionProfile := model.Permission{
			ID:   profilePermission,
			Code: "PM0001",
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
		}

		createOrder := uuid.MustParse("ab47d368-974c-4a99-aaba-8f0c2c853b55")
		createOrderPermission := model.Permission{
			ID:   createOrder,
			Code: "PM0002",
			Name: "create order",
			Path: "/create-order",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.New(),
					RoleID:         buyer,
					PermissionName: "create order",
					PermissionPath: "/create-order",
				},
			},
		}

		createBudidaya := uuid.MustParse("ab47d368-974c-4a99-aaba-8f0c2c853b55")
		createBudidayaPermission := model.Permission{
			ID:   createBudidaya,
			Code: "PM0003",
			Name: "create budidaya",
			Path: "/create-budidaya",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.New(),
					RoleID:         buyer,
					PermissionName: "create budidaya",
					PermissionPath: "/create-budidaya",
				},
			},
		}

		createPond := uuid.MustParse("6b4f540f-eaec-40b5-b54d-8cae44faf33b")
		createPondPermission := model.Permission{
			ID:   createPond,
			Code: "PM0004",
			Name: "create pond",
			Path: "/create-pond",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.New(),
					RoleID:         seller,
					PermissionName: "create pond",
					PermissionPath: "/create-pond",
				},
			},
		}

		permission = append(permission,
			permissionProfile,
			createOrderPermission,
			createBudidayaPermission,
			createPondPermission,
		)

		db.Save(&permission)
		return nil
	}
	return nil
}
