package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BILog struct {
	ClientId string `json:"clientId" validate:"required"`
	UseId    string `json:"user_id" validate:"required"`
	BIKey    string `json:"bi_key" validate:"required"`
	Value    string `json:"value" validate:"value"`
}

var sqlDb *gorm.DB

func init() {
	InitDB()
}

func InitDB() {
	// 创建 SQLite3 数据库连接
	db, err := gorm.Open(sqlite.Open("slg.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDb = db
	// 设置连接池大小
	//sqlDB, err := sqlDb.DB()
	//if err != nil {
	//	panic("failed to set database pool size")
	//}
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(100)

}

func GetSqlDb() *gorm.DB {
	return sqlDb
}
