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
		err := db.Debug().AutoMigrate(
			&Pond{},
			&User{},
			&Role{},
			&Permission{},
			&RolePermission{},
			&UserRole{},
			&UserPermission{},
			&Team{},
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
			&Banner{},
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
					RoleID:         seller,
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

		CreateFishSpecies := uuid.MustParse("54c6eb6b-4faa-5e0b-8360-f2674d6ce097")
		CreateFishSpeciesPermission := model.Permission{
			ID:   CreateFishSpecies,
			Code: "PM0007",
			Name: "create fish species",
			Path: "/create-fish-species",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("b5fecfec-1602-55fb-8e87-55da3bc6dbc1"),
					RoleID:         seller,
					PermissionName: "create fish species",
					PermissionPath: "/create-fish-species",
				},
				{
					ID:             uuid.MustParse("05cfddc4-e7e6-55a5-93dc-c6358c1a93b5"),
					RoleID:         admin,
					PermissionName: "create fish species",
					PermissionPath: "/create-fish-species",
				},
			},
		}

		createMultiplePriceList := uuid.MustParse("cf7813cd-f911-5f5e-a893-1f549b21f896")
		createMultiplePriceListPermission := model.Permission{
			ID:   createMultiplePriceList,
			Code: "PM0008",
			Name: "create multiple pricelist",
			Path: "/create-multiple-pricelist",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("d21c3d7a-f00d-5f70-aa29-21f5b1733bee"),
					RoleID:         seller,
					PermissionName: "create multiple pricelist",
					PermissionPath: "/create-multiple-pricelist",
				},
			},
		}

		listBudidayaSeller := uuid.MustParse("941787f8-55cc-597b-a858-432e12a55b99")
		listBudidayaSellerPermission := model.Permission{
			ID:   listBudidayaSeller,
			Code: "PM0009",
			Name: "list budidaya seller",
			Path: "/list-budidaya-seller",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("851f8565-32dd-57aa-84e1-3e9abf0fe24d"),
					RoleID:         seller,
					PermissionName: "list budidaya seller",
					PermissionPath: "/list-budidaya-seller",
				},
			},
		}

		listBudidaya := uuid.MustParse("20f07e9b-d285-5c29-a2e6-413153ea9cd8")
		listBudidayaPermission := model.Permission{
			ID:   listBudidaya,
			Code: "PM0010",
			Name: "list budidaya",
			Path: "/list-budidaya",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("550416e6-57ee-5872-b982-2ebdb1a47cc1"),
					RoleID:         admin,
					PermissionName: "list budidaya",
					PermissionPath: "/list-budidaya",
				}, {
					ID:             uuid.MustParse("b2ba22e5-34c0-56f7-a8ac-a337d5fa28c3"),
					RoleID:         buyer,
					PermissionName: "list budidaya",
					PermissionPath: "/list-budidaya",
				},
			},
		}

		updatePondStatus := uuid.MustParse("189c5cfd-a737-5fe9-9f3f-3654e666529b")
		updatePondStatusPermission := model.Permission{
			ID:   updatePondStatus,
			Code: "PM0011",
			Name: "update pond status",
			Path: "/update-pond-status",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("8c36c2e8-4e09-55e6-80ea-7d8e6d72c20b"),
					RoleID:         admin,
					PermissionName: "update pond status",
					PermissionPath: "/update-pond-status",
				},
			},
		}

		updatePond := uuid.MustParse("ae06c9a8-8f3a-59c7-a10e-57fe682d992d")
		updatePondPermission := model.Permission{
			ID:   updatePond,
			Code: "PM0012",
			Name: "update pond",
			Path: "/update-pond",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("84f3d099-a475-5f84-b3c1-55452a403dde"),
					RoleID:         seller,
					PermissionName: "update pond",
					PermissionPath: "/update-pond",
				},
			},
		}

		updateUser := uuid.MustParse("4017a3d6-cbbc-4b6a-9c01-58e1fe1e9e06")
		updateUserPermission := model.Permission{
			ID:   updateUser,
			Code: "PM0013",
			Name: "update user",
			Path: "/update-user",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("182df938-6615-410e-9bae-35736286dba6"),
					RoleID:         seller,
					PermissionName: "update user",
					PermissionPath: "/update-user",
				}, {
					ID:             uuid.MustParse("75742bec-e7d7-454c-93f8-1f9ca9904d92"),
					RoleID:         admin,
					PermissionName: "update user",
					PermissionPath: "/update-user",
				}, {
					ID:             uuid.MustParse("196f5fd3-8814-4d97-8ad1-dbc296009cf3"),
					RoleID:         buyer,
					PermissionName: "update user",
					PermissionPath: "/update-user",
				},
			},
		}

		createBanner := uuid.MustParse("9708d038-6be0-556e-9ffe-a2855178538a")
		createBannerPermission := model.Permission{
			ID:   createBanner,
			Code: "PM0014",
			Name: "create banner",
			Path: "/create-banner",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("2814db86-955b-55dd-9a20-d9d126d2389e"),
					RoleID:         admin,
					PermissionName: "create banner",
					PermissionPath: "/create-banner",
				},
			},
		}

		updateBanner := uuid.MustParse("e8d67381-841c-5c79-867d-fa270dcdede2")
		updateBannerPermission := model.Permission{
			ID:   updateBanner,
			Code: "PM0015",
			Name: "update banner",
			Path: "/update-banner",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("cf1181c4-f3e8-5594-84df-a79871cebaa7"),
					RoleID:         admin,
					PermissionName: "update banner",
					PermissionPath: "/update-banner",
				},
			},
		}

		updateCancelOrder := uuid.MustParse("60ca4c5b-9c16-53b4-b403-d0b7fadeb3b6")
		updateCancelOrderPermission := model.Permission{
			ID:   updateCancelOrder,
			Code: "PM0016",
			Name: "update cancel order",
			Path: "/update-order-cancel",
			RolePermission: []*model.RolePermission{{
				ID:             uuid.MustParse("bfc634db-bc90-51df-b8ef-b592675d9b9f"),
				RoleID:         buyer,
				PermissionName: "update cancel order",
				PermissionPath: "/update-order-cancel",
			}},
		}

		UpdateSuccessOrder := uuid.MustParse("177141be-21a6-5ea3-8b28-68ecf1487a07")
		UpdateSuccessOrderPermission := model.Permission{
			ID:   UpdateSuccessOrder,
			Code: "PM0017",
			Name: "update success order",
			Path: "/update-order-success",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("ce82bb38-b872-59b1-8b5b-0b67efb1a078"),
					RoleID:         seller,
					PermissionName: "update success order",
					PermissionPath: "/update-order-success",
				}, {
					ID:             uuid.MustParse("0ba03e3a-ba7b-5370-bcbf-d073713e88b0"),
					RoleID:         buyer,
					PermissionName: "update success order",
					PermissionPath: "/update-order-success",
				},
			},
		}

		getOrderCancel := uuid.MustParse("47eb02a8-0150-5351-a9ba-d2a48c020ec5")
		getOrderCancelPermission := model.Permission{
			ID:   getOrderCancel,
			Code: "PM0018",
			Name: "order cancel",
			Path: "/order-cancel",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("0a587772-c6d4-59d2-a382-6ac7291b8cae"),
					RoleID:         seller,
					PermissionName: "order cancel",
					PermissionPath: "/order-cancel",
				},
				{
					ID:             uuid.MustParse("9192bb3f-8116-5e84-9eaa-d1b09e5329a5"),
					RoleID:         buyer,
					PermissionName: "order cancel",
					PermissionPath: "/order-cancel",
				},
			},
		}

		getOrderSuccess := uuid.MustParse("7245c3d1-7f9e-51a3-990c-ced81ab3c0ab")
		getOrderSuccessPermission := model.Permission{
			ID:   getOrderSuccess,
			Code: "PM0019",
			Name: "order success",
			Path: "/order-success",
			RolePermission: []*model.RolePermission{
				{
					ID:             uuid.MustParse("4901dad9-4223-5b07-a880-bcd506ab63d0"),
					RoleID:         seller,
					PermissionName: "order success",
					PermissionPath: "/order-success",
				},
				{
					ID:             uuid.MustParse("77816b2e-3c92-5f79-afd1-9040a3c523b5"),
					RoleID:         buyer,
					PermissionName: "order successs",
					PermissionPath: "/order-success",
				},
			},
		}

		permission = append(permission,
			permissionProfile,
			createOrderPermission,
			createBudidayaPermission,
			createBudidayaPermission,
			createPondPermission,
			getPondSellerPermission,
			getOrderPermission,
			CreateFishSpeciesPermission,
			createMultiplePriceListPermission,
			listBudidayaPermission,
			listBudidayaSellerPermission,
			updatePondStatusPermission,
			updatePondPermission,
			updateUserPermission,
			createBannerPermission,
			updateBannerPermission,
			updateCancelOrderPermission,
			UpdateSuccessOrderPermission,
			getOrderCancelPermission,
			getOrderSuccessPermission,
		)

		db.Save(&permission)
		return nil
	}
	return nil
}
