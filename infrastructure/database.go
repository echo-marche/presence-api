package infrastructure

import (
	"github.com/alduh-co/presence-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDb() *gorm.DB {
	// TODO SounDiverseDBに接続→パラメータにより接続先を変更
	dbConfig := config.InitSounDiverseDBConfig()

	db, err := gorm.Open("mysql", dbConfig.User+":"+dbConfig.Password+"@("+dbConfig.Host+":"+dbConfig.Port+")/"+dbConfig.Name+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database.")
	}
	db.LogMode(true)

	return db
}
