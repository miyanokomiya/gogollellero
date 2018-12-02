package models

// User ユーザー
type User struct {
	Model
	Name string `gorm:"size:255;unique_index"`
}

// Create 作成
func Create(user *User) *User {
	db := gormConnect()
	db.Create(user)
	created := User{}
	db.First(&created, "Name = ?", user.Name)
	return &created
}

// Delete 削除
func Delete(id uint) {
	db := gormConnect()
	user := User{}
	user.ID = id
	db.Delete(&user)
}
