package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "faftechm_baiturrahman:@Labithotspot123@/faftechm_simcovid?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "root@/simcovid?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	//db.AutoMigrate(structs.User{}, structs.News{}, structs.Hoax{}, structs.Protokol{}, structs.EducationCategory{},
	//	structs.Education{}, structs.QnA{})
	return db
}
