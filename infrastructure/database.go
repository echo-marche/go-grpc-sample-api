package infrastructure

import (
	"github.com/echo-marche/hack-tech-tips-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDb() *gorm.DB {
	dbConfig := config.InitHackTechTipsDBConfig()
	db, err := gorm.Open("mysql", dbConfig.User+":"+dbConfig.Password+"@("+dbConfig.Host+":"+dbConfig.Port+")/"+dbConfig.Name+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database.")
	}
	db.LogMode(true)

	return db
}
