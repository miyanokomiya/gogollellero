package models

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	yaml "gopkg.in/yaml.v2"
)

// DB DBインスタンス
var DB *gorm.DB

func readConfig() string {
	yml, err := ioutil.ReadFile("../../../configs/db.yml")
	if err != nil {
		panic(err)
	}
	t := make(map[interface{}]interface{})
	_ = yaml.Unmarshal([]byte(yml), &t)
	conn := t[os.Getenv("GO_ENV")].(map[interface{}]interface{})
	protocol := t["protocol"].(string)
	return conn["user"].(string) + ":" + conn["password"].(string) + "@" + protocol + "/" + conn["db"].(string) + "?charset=utf8&parseTime=True"
}

// GormOpen 接続
func GormOpen() {
	if DB != nil {
		return
	}
	var err error
	DB, err = gorm.Open("mysql", readConfig())
	if err != nil {
		panic(err)
	}
}

// GormClose 切断
func GormClose() {
	if DB != nil {
		DB.Close()
	}
	DB = nil
}

// Model 基底モデル
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
