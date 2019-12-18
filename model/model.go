package model

import (
	"favorite-jobs/config"
	"favorite-jobs/log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var (
	once sync.Once
	DB *gorm.DB
	err error
)

func InitDB() {
	once.Do(func() {
		connStr := fmt.Sprintf(
			"%s:%s8@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config.DB.UserName,
			config.Config.DB.Password,
			config.Config.DB.Addr,
			config.Config.DB.Port,
			config.Config.DB.Database,
			)
		DB, err = gorm.Open("mysql", connStr)
		if err != nil {
			log.ZapLog.Infow("数据库初始化失败", "错误原因", err)
			return
		}
		//DB.LogMode(true)
	})
}


