package models

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestCreate(t *testing.T) {
	GormOpen()
	user := User{Name: "test_abcd", Password: "abcdabcd"}
	err := user.Create()
	if err != nil {
		t.Fatal("failed test", err)
	}
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
	user := User{Name: "test_abcd", Password: "abcdabcd"}
	DB.Create(&user)
	user2 := User{Name: "test_dddd", Password: "abcdabcd"}
	DB.Create(&user2)
	defer DB.Delete(&user)
	defer DB.Delete(&user2)
	// id検索
	read := User{}
	read.ID = user2.ID
	read.Read()
	if read.Name != "test_dddd" {
		t.Fatal("failed read", read)
	}
	// name検索
	read2 := User{}
	read2.Name = "test_dddd"
	read2.Read()
	if read2.Name != "test_dddd" {
		t.Fatal("failed read", read2)
	}
	// 検索エラー
	read3 := User{}
	if read3.Read() == nil {
		t.Fatal("invalid read")
	}
}

func TestUpdate(t *testing.T) {
	GormOpen()
	// 作成
	user := User{Name: "test_abcd", Password: "abcdabcd"}
	DB.Create(&user)
	defer DB.Delete(&user)
	created := User{}
	DB.First(&created, "Name = ?", user.Name)

	// 更新
	created.Name = "updated"
	if err := created.Update(); err != nil {
		t.Fatal("failed update", err)
	}
	if created.Name != "updated" {
		t.Fatal("failed update", created)
	}
}

func TestUpdate2(t *testing.T) {
	GormOpen()
	// 作成
	user := User{Name: "test_abcd", Password: "abcdabcd"}
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
	userListWrapper(1, func(users Users) {
		user := users[0]
		user.Delete()
		deleted := User{}
		DB.First(&deleted, user.ID)
		if deleted.ID != 0 {
			t.Fatal("failed delete", deleted)
		}
	})
}

func TestIndex(t *testing.T) {
	userListWrapper(3, func(_ Users) {
		users := Users{}
		users.Index(nil)
		if len(users) != 3 {
			t.Fatal("failed test", len(users))
		}
	})
}

func TestBatchDelete(t *testing.T) {
	userListWrapper(3, func(created Users) {
		BatchDelete([]int{created[0].ID, created[1].ID, created[2].ID})
		users := Users{}
		users.Index(nil)
		if len(users) != 0 {
			t.Fatal("failed delete", len(users))
		}
	})
}

func TestAuthenticate(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal("failed hash")
	}
	user := User{
		Name:     "name",
		Password: string(hash),
	}
	if !user.Authenticate("password") {
		t.Fatal("failed auth when valid password")
	}
	if user.Authenticate("password1") {
		t.Fatal("failed auth when invalid password")
	}
}

func TestSetPassword(t *testing.T) {
	user := User{}
	user.SetPassword("password")
	if user.Password == "password" {
		t.Fatal("failed set hashed password")
	}
	if !user.Authenticate("password") {
		t.Fatal("failed set hashed password")
	}
}

func TestSetPassword2(t *testing.T) {
	user := User{}
	err := user.SetPassword("")
	if err == nil {
		t.Fatal("empty password allowed")
	}
}
