package db

import (
	"testing"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func init() {
	InitDB()
}

func TestSql(t *testing.T) {
	db := GetSqlDb()
	err1 := db.AutoMigrate(&Product{})
	if err1 != nil {
		t.Log(err1.Error())
		return
	}

	db.Create(&Product{Code: "D42", Price: 100})
}
