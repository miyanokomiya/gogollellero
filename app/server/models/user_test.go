package models

import (
	"testing"
)

func TestCreate(t *testing.T) {
	GormOpen()
	// 作成
	user := User{Name: "test_abcd"}
	user.Create()
	defer DB.Delete(&user)
	if user.Name != "test_abcd" {
		t.Fatal("failed test", user)
	}
	if user.ID == 0 {
		t.Fatal("failed test", user)
	}
}

func TestRead(t *testing.T) {
	GormOpen()
	// 作成
	user := User{Name: "test_abcd"}
	DB.Create(&user)
	defer DB.Delete(&user)
	read := User{}
	read.ID = user.ID
	read.Read()
	if read.Name != "test_abcd" {
		t.Fatal("failed read", read)
	}
}

func TestUpdate(t *testing.T) {
	GormOpen()
	// 作成
	user := User{Name: "test_abcd"}
	DB.Create(&user)
	defer DB.Delete(&user)
	created := User{}
	DB.First(&created, "Name = ?", user.Name)

	// 更新
	update := User{}
	update.ID = created.ID
	update.Name = "updated"
	update.Update()
	if update.Name != "updated" {
		t.Fatal("failed update", update)
	}
}

func TestUpdate2(t *testing.T) {
	GormOpen()
	// 作成
	user := User{Name: "test_abcd"}
	DB.Create(&user)
	defer DB.Delete(&user)
	created := User{}
	DB.First(&created, "Name = ?", user.Name)

	// 更新
	created.Update()
	if created.Name != "test_abcd" {
		t.Fatal("failed update", created)
	}
}

func TestDelete(t *testing.T) {
	GormOpen()
	// 作成
	user := User{Name: "test_abcd"}
	DB.Create(&user)
	defer DB.Delete(&user)

	// 削除
	user.Delete()
	deleted := User{}
	DB.First(&deleted, user.ID)
	if deleted.ID != 0 {
		t.Fatal("failed delete", deleted)
	}
}

func TestIndex(t *testing.T) {
	GormOpen()
	user1 := User{Name: "user1"}
	DB.Create(&user1)
	defer DB.Delete(&user1)
	user2 := User{Name: "user2"}
	DB.Create(&user2)
	defer DB.Delete(&user2)
	user3 := User{Name: "user3"}
	DB.Create(&user3)
	defer DB.Delete(&user3)

	users := Users{}
	users.Index()
	if len(users) != 3 {
		t.Fatal("failed delete", len(users))
	}
	if users[0].Name != "user1" {
		t.Fatal("failed delete", users[0])
	}
}

func TestBatchDelete(t *testing.T) {
	GormOpen()
	user1 := User{Name: "user1"}
	DB.Create(&user1)
	defer DB.Delete(&user1)
	user2 := User{Name: "user2"}
	DB.Create(&user2)
	defer DB.Delete(&user2)
	user3 := User{Name: "user3"}
	DB.Create(&user3)
	defer DB.Delete(&user3)

	BatchDelete([]uint{user1.ID, user2.ID, user3.ID})
	users := Users{}
	users.Index()
	if len(users) != 0 {
		t.Fatal("failed delete", len(users))
	}
}
