package migrations_test

import (
	"log"
	"testing"

	"github.com/e-fish/api/migrations"
	"github.com/e-fish/api/pkg/common/helper/config"
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

	err = migrations.Migrate(con, "initial-data-model")
	if err != nil {
		log.Println("error: ", err)
		return
	}

}

func Test_Migrate_RegionData(t *testing.T) {
	logger.SetupLogger("true")

	conMysql, err := orm.CreateConnetionDB(migrations.GetConfig().DbConfig)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	conPostgress, err := orm.CreateConnetionDB(config.DbConfig{
		Driver:   "postgres",
		Host:     "mosleim.com",
		User:     "palen_admin",
		Password: "adminpalen",
		Database: "palen",
		Port:     "5434",
	})
	if err != nil {
		log.Println("error: ", err)
		return
	}

	Country := []migrations.Country{}

	err = conPostgress.Preload("ListProvince.ListCity.ListDistrict").Find(&Country).Error
	if err != nil {
		log.Println("err", err)
		return
	}

	dbTxn := conMysql.Begin()

	for _, v := range Country {
		err = dbTxn.Omit("ListProvince.ListCity.ListDistrict").Create(&v).Error
		if err != nil {
			dbTxn.Rollback()
			log.Println("error create country")
			return
		}
		for _, w := range v.ListProvince {
			err = dbTxn.Omit("ListCity.ListDistrict").Save(&w).Error
			if err != nil {
				dbTxn.Rollback()
				log.Println("error create province")
				return
			}
			for _, x := range w.ListCity {
				err = dbTxn.Omit("ListDistrict").Save(&x).Error
				if err != nil {
					dbTxn.Rollback()
					log.Println("error create city")
					return
				}
				for _, y := range x.ListDistrict {
					err = dbTxn.Save(&y).Error
					if err != nil {
						dbTxn.Rollback()
						log.Println("error create district")
						return
					}
				}
			}
		}
	}

	if err == nil {
		dbTxn.Commit()
	}
}
