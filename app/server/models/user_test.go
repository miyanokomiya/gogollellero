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
