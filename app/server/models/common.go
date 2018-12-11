package models

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql用のimport
	"github.com/miyanokomiya/gogollellero/app/server/assets"
	v "gopkg.in/go-playground/validator.v8"
	yaml "gopkg.in/yaml.v2"
)

// DB DBインスタンス
var DB *gorm.DB
var validator *v.Validate

// GetValidator バリデーションインスタンス取得
func GetValidator() *v.Validate {
	if validator != nil {
		return validator
	}
	validator = v.New(&v.Config{TagName: "binding"})
	return validator
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

// Tx トランザクションラッパー
func Tx(fn func(*gorm.DB) error) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// Model 基底モデル
type Model struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Pagination ページネーション情報
type Pagination struct {
	Page    int
	Limit   int
	OrderBy string
	Keyword string
}

func paginate(db *gorm.DB, pagination *Pagination) *gorm.DB {
	if pagination == nil {
		return db
	}
	if pagination.OrderBy != "" {
		db = db.Order(pagination.OrderBy)
	}
	if pagination.Page > 0 && pagination.Limit > 1 {
		db = db.Offset(pagination.Limit * (pagination.Page - 1)).Limit(pagination.Limit)
	}
	return db
}

func parseYml(file http.File) (map[interface{}]interface{}, error) {
	by := new(bytes.Buffer)
	io.Copy(by, file)
	buf := by.Bytes()
	t := make(map[interface{}]interface{})
	err := yaml.Unmarshal(buf, &t)
	return t, err
}

func readConfig() string {
	file, err := assets.Configs.Open("/configs/db.yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	t, err := parseYml(file)
	if err != nil {
		panic(err)
	}
	conn := t[os.Getenv("GO_ENV")].(map[interface{}]interface{})
	protocol := t["protocol"].(string)
	return conn["user"].(string) + ":" + conn["password"].(string) + "@" + protocol + "/" + conn["db"].(string) + "?charset=utf8&parseTime=True"
}
