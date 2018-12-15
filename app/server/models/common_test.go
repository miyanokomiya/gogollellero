package models

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestGormOpen(t *testing.T) {
	GormOpen()
	if DB == nil {
		t.Fatal("failed test")
	}
}

func TestGormClose(t *testing.T) {
	GormOpen()
	if DB == nil {
		t.Fatal("failed test")
	}
	GormClose()
	if DB != nil {
		t.Fatal("failed test")
	}
}

func TestTxSuccess(t *testing.T) {
	GormOpen()
	user1 := User{Name: "user1", Password: "abcdabcd"}
	Tx(func(db *gorm.DB) error {
		return db.Create(&user1).Error
	})
	defer DB.Delete(&user1)
	created := User{}
	created.ID = user1.ID
	if err := created.Read(); err != nil {
		t.Fatal("failed test", err)
	}
}

func TestTxRollback(t *testing.T) {
	GormOpen()
	user1 := User{Name: "user1", Password: "abcdabcd"}
	user2 := User{Name: "user1", Password: "abcdabcd"}
	err := Tx(func(db *gorm.DB) error {
		if err := db.Create(&user1).Error; err != nil {
			t.Fatal("failed test", err)
		}
		return db.Create(&user2).Error
	})
	defer DB.Delete(&user1)
	if err == nil {
		t.Fatal("failed test", err)
	}
	created := User{}
	created.ID = user1.ID
	if err := created.Read(); err == nil {
		t.Fatal("failed test", err)
	}
}

func userListWrapper(count int, fn func(Users)) {
	GormOpen()
	var users Users
	for i := 0; i < count; i++ {
		user := User{Name: fmt.Sprintf("user_%d", i), Password: "abcdabcd"}
		DB.Create(&user)
		defer DB.Delete(&user)
		users = append(users, user)
	}
	fn(users)
}

func tagListWrapper(count int, fn func(Tags)) {
	GormOpen()
	var tags Tags
	for i := 0; i < count; i++ {
		tag := Tag{Title: fmt.Sprintf("tag_%d", i)}
		DB.Create(&tag)
		defer DB.Delete(&tag)
		tags = append(tags, tag)
	}
	fn(tags)
}

func postListWrapper(count int, fn func(Posts)) {
	GormOpen()
	var posts Posts
	tag := Tag{Title: "tag"}
	DB.Create(&tag)
	defer DB.Delete(&tag)
	for i := 0; i < count; i++ {
		user := User{
			Name:     fmt.Sprintf("user_%d", i),
			Password: "password",
		}
		DB.Create(&user)
		defer DB.Delete(&user)
		var tags Tags
		tags = append(tags, tag)
		post := Post{
			UserID:   user.ID,
			Title:    fmt.Sprintf("title_%d", i),
			Problem:  "problem",
			Solution: "solution",
			Lesson:   "lesson",
			Tags:     tags,
		}
		post.Create()
		defer DB.Delete(post.PostParent)
		posts = append(posts, post)
	}
	fn(posts)
}

func TestPaginate(t *testing.T) {
	userListWrapper(10, func(_ Users) {
		users := Users{}

		users.Index(&Pagination{
			Page:    1,
			Limit:   3,
			OrderBy: "id asc",
		})
		if len(users) != 3 {
			t.Fatal("failed test", len(users))
		}
		if users[0].Name != "user_0" {
			t.Fatal("failed test", users[0])
		}

		users.Index(&Pagination{
			Page:    2,
			Limit:   3,
			OrderBy: "id asc",
		})
		if len(users) != 3 {
			t.Fatal("failed test", len(users))
		}
		if users[0].Name != "user_3" {
			t.Fatal("failed test", users[0])
		}

		users.Index(&Pagination{
			Page:    4,
			Limit:   3,
			OrderBy: "id asc",
		})
		if len(users) != 1 {
			t.Fatal("failed test", len(users))
		}
		if users[0].Name != "user_9" {
			t.Fatal("failed test", users[0])
		}

		users.Index(&Pagination{
			Page:    5,
			Limit:   3,
			OrderBy: "id asc",
		})
		if len(users) != 0 {
			t.Fatal("failed test", len(users))
		}

		// 全取得
		users.Index(&Pagination{
			Page:    0,
			Limit:   0,
			OrderBy: "id asc",
		})
		if len(users) != 10 {
			t.Fatal("failed test", len(users))
		}
		users.Index(nil)
		if len(users) != 10 {
			t.Fatal("failed test", len(users))
		}

		// ソート
		users.Index(&Pagination{
			Page:    2,
			Limit:   3,
			OrderBy: "name desc",
		})
		if users[0].Name != "user_6" {
			t.Fatal("failed test", users[0])
		}
	})
}
