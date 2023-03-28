package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ginkwok/artihub/model"
)

func GetDB(host string, port string, user string, pass string, database string) *gorm.DB {
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect mysql db:" + err.Error())
	}

	err = db.AutoMigrate(&model.Article{})
	if err != nil {
		panic(err)
	}

	return db
}
