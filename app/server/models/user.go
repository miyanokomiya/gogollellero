package models

// User ユーザー
type User struct {
	Model
	Name string `gorm:"size:255;unique_index"`
}

// Create 作成
func (user *User) Create() {
	db := gormConnect()
	db.Create(user)
}

// Read 読込
func (user *User) Read() {
	db := gormConnect()
	db.First(user)
}

// Update 更新
func (user *User) Update() {
	db := gormConnect()
	db.Save(user)
}

// Delete 削除
func (user *User) Delete() {
	db := gormConnect()
	db.Delete(user)
}
