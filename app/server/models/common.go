package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbms = "mysql"
var user = "miyanokomiya"
var pass = "miyanokomiya"
var protocol = "tcp(gogollellero_db:3306)"
var dbname = "gogollellero"
var connect = user + ":" + pass + "@" + protocol + "/" + dbname

func gormConnect() *gorm.DB {
	db, err := gorm.Open(dbms, connect)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// Model 基底モデル
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
