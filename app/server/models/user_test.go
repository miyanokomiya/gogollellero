package models

import (
	"testing"
)

func TestCreate(t *testing.T) {
	// 作成
	user := User{Name: "test_abcd"}
	crated := Create(&user)
	if crated.Name != "test_abcd" {
		t.Fatal("failed test", crated.ID)
	}

	// 削除
	db := gormConnect()
	db.Delete(&crated)
}

func TestDelete(t *testing.T) {
	// 作成
	user := User{Name: "test_abcd"}
	db := gormConnect()
	db.Create(&user)
	created := User{}
	db.First(&created, "Name = ?", user.Name)
	if created.ID == 0 {
		t.Fatal("failed create")
	}

	// 削除
	Delete(created.ID)
	deleted := User{}
	db.First(&deleted, created.ID)
	if deleted.ID != 0 {
		t.Fatal("failed delete", deleted)
	}
}
