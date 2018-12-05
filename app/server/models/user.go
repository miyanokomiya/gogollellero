package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// User ユーザー
type User struct {
	Model
	Name     string `gorm:"size:255;unique_index"`
	Password string `gorm:"size:255"`
}

// Users ユーザー一覧
type Users []User

// Create 作成
func (user *User) Create() error {
	return DB.Create(user).Error
}

// Read 読込
func (user *User) Read() error {
	return DB.First(user).Error
}

// Update 更新
func (user *User) Update() error {
	return DB.Save(user).Error
}

// Delete 削除
func (user *User) Delete() error {
	return DB.Delete(user).Error
}

// Index 一覧
func (users *Users) Index(pagination *Pagination) error {
	return paginate(DB, pagination).Find(users).Error
}

// BatchDelete 一覧削除
func BatchDelete(ids []uint) error {
	return DB.Where(ids).Delete(User{}).Error
}

// Authenticate 認証
func (user *User) Authenticate(password string) bool {
	user.Read()
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("Failed to authenticate", err)
		return false
	}
	return true
}

// SetPassword パスワードをハッシュ化してセット
func (user *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}
