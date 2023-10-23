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

var sqlDb


func InitDB() {
	// 创建 SQLite3 数据库连接
	db, err := gorm.Open(sqlite.Open("slg.db"), &gorm.Config{
		// 开启 WAL 模式
		DSN: "mode=wal",
	})
	if err != nil {
		panic("failed to connect database")
	}
	// 设置连接池大小
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to set database pool size")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// 自动迁移 User 模型对应的表
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate table")
	}
}

