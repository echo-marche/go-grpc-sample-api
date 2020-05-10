package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/echo-marche/hack-tech-tips-api/config"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
)

func NewDbMap() *gorp.DbMap {
	dbConfig := config.InitHackTechTipsDBConfig()

	db, err := sql.Open("mysql", dbConfig.User+":"+dbConfig.Password+"@("+dbConfig.Host+":"+dbConfig.Port+")/"+dbConfig.Name+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database.")
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	if config.IsDev() {
		dbMap.TraceOn("[gorp]", log.New(os.Stdout, "", log.Lmicroseconds))
	}
	return dbMap
}
