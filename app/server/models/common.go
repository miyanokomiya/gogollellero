package models

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	yaml "gopkg.in/yaml.v2"
)

var db *gorm.DB

func gormConnect() *gorm.DB {
	if db != nil {
		return db
	}
	yml, err := ioutil.ReadFile("../../../configs/db.yml")
	if err != nil {
		panic(err)
	}
	t := make(map[interface{}]interface{})
	_ = yaml.Unmarshal([]byte(yml), &t)
	conn := t[os.Getenv("GO_ENV")].(map[interface{}]interface{})
	protocol := t["protocol"].(string)
	db, err = gorm.Open("mysql", conn["user"].(string)+":"+conn["password"].(string)+"@"+protocol+"/"+conn["db"].(string)+"?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	return db
}

// Model 基底モデル
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
