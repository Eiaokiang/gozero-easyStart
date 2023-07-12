package store

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDbStore() *gorm.DB {
	user := viper.GetString("database.mysql.user")
	password := viper.GetString("database.mysql.password")
	host := viper.GetString("database.mysql.host")
	port := viper.GetString("database.mysql.port")
	dbName := viper.GetString("database.mysql.dbName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db

}
