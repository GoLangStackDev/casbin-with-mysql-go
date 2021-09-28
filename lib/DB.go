package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Gorm *gorm.DB
func InitDB()  {
	Gorm = InitGorm()
}
func InitGorm() *gorm.DB {
	dsn:=InitConfig().Mysql.DNS
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB,err:=db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	return db
}