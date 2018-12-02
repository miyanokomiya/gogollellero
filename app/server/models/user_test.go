package models

import (
	"testing"
)

func TestCreate(t *testing.T) {
	db := gormConnect()
	// 作成
	user := User{Name: "test_abcd"}
	user.Create()
	defer db.Delete(&user)
	if user.Name != "test_abcd" {
		t.Fatal("failed test", user)
	}
	if user.ID == 0 {
		t.Fatal("failed test", user)
	}
}

func TestRead(t *testing.T) {
	db := gormConnect()
	// 作成
	user := User{Name: "test_abcd"}
	db.Create(&user)
	defer db.Delete(&user)
	read := User{}
	read.ID = user.ID
	read.Read()
	if read.Name != "test_abcd" {
		t.Fatal("failed read", read)
	}
}

func TestUpdate(t *testing.T) {
	db := gormConnect()
	// 作成
	user := User{Name: "test_abcd"}
	db.Create(&user)
	defer db.Delete(&user)
	created := User{}
	db.First(&created, "Name = ?", user.Name)

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
	db := gormConnect()
	// 作成
	user := User{Name: "test_abcd"}
	db.Create(&user)
	defer db.Delete(&user)
	created := User{}
	db.First(&created, "Name = ?", user.Name)

	// 更新
	created.Update()
	if created.Name != "test_abcd" {
		t.Fatal("failed update", created)
	}
}

func TestDelete(t *testing.T) {
	db := gormConnect()
	// 作成
	user := User{Name: "test_abcd"}
	db.Create(&user)
	defer db.Delete(&user)

	// 削除
	user.Delete()
	deleted := User{}
	db.First(&deleted, user.ID)
	if deleted.ID != 0 {
		t.Fatal("failed delete", deleted)
	}
}
