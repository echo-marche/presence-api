package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/echo-marche/presence-api/config"
	"github.com/echo-marche/presence-api/models"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
)

func NewDbMap() *gorp.DbMap {
	// DBに接続→パラメータにより接続先を変更
	dbConfig := config.DBConfig{}
	switch config.GetEnv("SYSTEM_CODE") {
	case "hack-tech-tips":
		dbConfig = config.InitHackTechTipsDBConfig()
	default:
		panic("db config is nothing!")
	}

	db, err := sql.Open("mysql", dbConfig.User+":"+dbConfig.Password+"@("+dbConfig.Host+":"+dbConfig.Port+")/"+dbConfig.Name+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database.")
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	dbMap.AddTableWithName(models.User{}, "users").SetKeys(true, "Id")
	if config.IsDev() {
		dbMap.TraceOn("[gorp]", log.New(os.Stdout, "", log.Lmicroseconds))
	}
	return dbMap
}
