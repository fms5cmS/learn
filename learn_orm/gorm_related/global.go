package gorm_related

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

var (
	db   *gorm.DB
	once = new(sync.Once)
)

func init() {
	once.Do(func() {
		var (
			host, user, password, dbname, url string
			err error
		)
		viper.SetConfigName("conf")
		viper.AddConfigPath("./")
		if err := viper.ReadInConfig(); err != nil {
			logrus.Errorf("load config error: %v", err)
		}
		host = viper.GetString("mysql.host")
		user = viper.GetString("mysql.user")
		password = viper.GetString("mysql.password")
		dbname = viper.GetString("mysql.dbname")
		url = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbname)
		db, err = gorm.Open("mysql", url)
		if err != nil {
			panic("failed to connect database")
		}
		// 开启日志
		db.LogMode(true)
		// 全局禁用表名复数,如果设置为 true, `User` 的默认表名为 `user`,否则表名就是 `users`
		db.SingularTable(true)
		// 注册回调函数
		db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
		db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
		db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	})
}
