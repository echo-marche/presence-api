package infrastructure

import (
	"github.com/echo-marche/presence-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDb() *gorm.DB {
	// DBに接続→パラメータにより接続先を変更
	dbConfig := config.DBConfig{}
	switch config.GetEnv("SYSTEM_CODE") {
	case "niche-farm":
		dbConfig = config.InitNicheFarmDBConfig()
	default:
		panic("db config is nothing!")
	}

	db, err := gorm.Open("mysql", dbConfig.User+":"+dbConfig.Password+"@("+dbConfig.Host+":"+dbConfig.Port+")/"+dbConfig.Name+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database.")
	}
	db.LogMode(true)

	return db
}
