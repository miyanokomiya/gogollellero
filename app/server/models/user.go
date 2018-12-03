package models

// User ユーザー
type User struct {
	Model
	Name string `gorm:"size:255;unique_index"`
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

func (users *Users) Read() error {
	return DB.Order("id asc").Find(users).Error
}
