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
	case "update-user-data":

		err := db.Migrator().DropTable(&User{})
		if err != nil {
			return err
		}
		err = db.Migrator().DropTable(&Pond{})

		return err

	case "update-pond-data":
		dbTxn := db.Begin()
		err := dbTxn.Migrator().DropTable(&Pond{})
		defer func() {
			if err != nil {
				dbTxn.Rollback()
				return
			}
			dbTxn.Commit()
		}()
		if err != nil {
			return err
		}

		err = dbTxn.AutoMigrate(&Pond{})
		if err != nil {
			return err
		}
		return nil
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
					ID:             uuid.MustParse("548a9a4a-8fb2-5c6a-b357-691a90b208c4"),
					RoleID:         buyer,
					PermissionID:   profilePermission,
					PermissionName: "profile",
					PermissionPath: "/profile",
				},
				{
					ID:             uuid.MustParse("5bf8249a-6593-564e-b926-f07264e23fb0"),
					RoleID:         seller,
					PermissionID:   profilePermission,
					PermissionName: "profile",
					PermissionPath: "/profile",
				},
				{
					ID:             uuid.MustParse("b8a16a74-16b8-5b7c-bb1a-971beb7f942d"),
					RoleID:         admin,
					PermissionID:   profilePermission,
					PermissionName: "profile",
					PermissionPath: "/profile",
				},
			},
		}

		createOrder := uuid.MustParse("0774379d-49aa-49f3-b62a-41be8c6d61aa")
		createOrderPermission := model.Permission{
			ID:   createOrder,
			Code: "PM0002",
			Name: "create order",
			Path: "/create-order",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("d5cfaaef-d346-5e13-af7b-d22c429e5022"),
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
					ID:             uuid.MustParse("33120128-4910-54bd-8a8e-0f3ab0db8650"),
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
					ID:             uuid.MustParse("f1aeb356-4ff4-5590-a8b2-3ce77a7d3425"),
					RoleID:         seller,
					PermissionName: "create pond",
					PermissionPath: "/create-pond",
				},
			},
		}

		getPondSeller := uuid.MustParse("a2decae5-d84d-4073-aaf9-dc84d8ed8c33")
		getPondSellerPermission := model.Permission{
			ID:   getPondSeller,
			Code: "PM0005",
			Name: "pond",
			Path: "/pond",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("71927c27-902e-500c-a603-b90a373bddee"),
					RoleID:         seller,
					PermissionName: "pond",
					PermissionPath: "/pond",
				},
			},
		}

		getOrder := uuid.MustParse("aa19cc95-94da-56a2-9b46-5cfdb42e3985")
		getOrderPermission := model.Permission{
			ID:   getOrder,
			Code: "PM0006",
			Name: "order",
			Path: "/order",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("4fc8fc64-382d-5c6e-89c8-d2032c01cfad"),
					RoleID:         seller,
					PermissionName: "order",
					PermissionPath: "/order",
				},
				{
					ID:             uuid.MustParse("6c776404-4df1-5957-8b49-f2764fa60b7e"),
					RoleID:         buyer,
					PermissionName: "order",
					PermissionPath: "/order",
				},
			},
		}

		permission = append(permission,
			permissionProfile,
			createOrderPermission,
			createBudidayaPermission,
			createPondPermission,
			getPondSellerPermission,
			getOrderPermission,
		)

		db.Save(&permission)
		return nil
	}
	return nil
}
