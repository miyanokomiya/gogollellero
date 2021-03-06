package models

import (
	"fmt"
	"testing"
)

func TestCreatePost(t *testing.T) {
	GormOpen()
	user := User{
		Name:     "username",
		Password: "password",
	}
	if err := user.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer user.Delete()
	post := Post{
		UserID:   user.ID,
		Title:    "title",
		Problem:  "problem",
		Solution: "solution",
		Lesson:   "lesson",
	}
	if err := post.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer DB.Delete(post.PostParent)
	if post.Title != "title" {
		t.Fatal("failed test", post)
	}
	if post.ID == 0 {
		t.Fatal("failed test", post)
	}
	if post.Type != Draft {
		t.Fatal("failed test", post)
	}
}

func TestReadPost(t *testing.T) {
	GormOpen()
	user := User{
		Name:     "username",
		Password: "password",
	}
	if err := user.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer user.Delete()
	post := Post{
		UserID:   user.ID,
		Title:    "title",
		Problem:  "problem",
		Solution: "solution",
		Lesson:   "lesson",
	}
	if err := post.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer DB.Delete(post.PostParent)
	// id検索
	read := Post{}
	read.ID = post.ID
	read.Read()
	if read.Title != "title" {
		t.Fatal("failed read", read)
	}
	// PostParent付属
	if read.PostParent == nil || read.PostParent.ID == 0 {
		t.Fatal("failed read", read)
	}
	// 検索エラー
	read3 := Post{}
	if read3.Read() == nil {
		t.Fatal("invalid read", read3)
	}
}

func TestUpdatePost(t *testing.T) {
	GormOpen()
	user := User{
		Name:     "username",
		Password: "password",
	}
	if err := user.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer user.Delete()
	tags, _ := CreateTagsIfNotExist([]string{"a", "b", "c"})
	for _, tag := range tags {
		defer DB.Delete(&tag)
	}
	post := Post{
		UserID:   user.ID,
		Title:    "title",
		Problem:  "problem",
		Solution: "solution",
		Lesson:   "lesson",
		Tags:     tags,
	}
	if err := post.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer DB.Delete(post.PostParent)
	created := Post{}
	DB.First(&created, "ID = ?", post.ID)

	// 更新
	tags2, _ := CreateTagsIfNotExist([]string{"b", "c"})
	for _, tag := range tags2 {
		defer DB.Delete(&tag)
	}
	created.Title = "new_title"
	created.Tags = tags2
	if err := created.Update(); err != nil {
		t.Fatal("failed update", err)
	}
	updated := Post{}
	updated.ID = created.ID
	if err := updated.Read(); err != nil {
		t.Fatal("failed update", updated)
	}
	if updated.Title != "new_title" {
		t.Fatal("failed update", updated)
	}
	// タグ参照が置き換えられていること
	if len(updated.Tags) != 2 {
		t.Fatal("failed update", updated)
	}
	if updated.Tags[0].Title != "b" {
		t.Fatal("failed update", updated)
	}
	if updated.Tags[1].Title != "c" {
		t.Fatal("failed update", updated)
	}
}

func TestUpdatePost2(t *testing.T) {
	GormOpen()
	user := User{
		Name:     "username",
		Password: "password",
	}
	if err := user.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer user.Delete()
	post := Post{
		UserID:   user.ID,
		Title:    "title",
		Problem:  "problem",
		Solution: "solution",
		Lesson:   "lesson",
	}
	if err := post.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer DB.Delete(post.PostParent)
	created := Post{}
	DB.First(&created, "ID = ?", post.ID)

	// 更新
	created.Update()
	if created.Title != "title" {
		t.Fatal("failed update", created)
	}
}

func TestUpdatePostOnlyDraft(t *testing.T) {
	GormOpen()
	user := User{
		Name:     "username",
		Password: "password",
	}
	if err := user.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer user.Delete()
	post := Post{
		UserID: user.ID,
		Title:  "title",
	}
	if err := post.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer DB.Delete(post.PostParent)
	created := Post{}
	DB.First(&created, "ID = ?", post.ID)

	// 更新
	created.Type = Published
	if err := created.Update(); err == nil {
		t.Fatal("failed update", created)
	}
}

func TestDeletePost(t *testing.T) {
	postListWrapper(1, func(posts Posts) {
		post := posts[0]
		post.Delete()
		deleted := Post{}
		DB.First(&deleted, post.ID)
		if deleted.ID != 0 {
			t.Fatal("failed delete", deleted)
		}
	})
}

func TestPublish(t *testing.T) {
	postListWrapper(1, func(posts Posts) {
		post := posts[0]
		if published, err := post.Publish(); err != nil {
			t.Fatal("failed publish", err)
		} else {
			if published.ID == post.ID {
				t.Fatal("failed publish", published)
			}
		}
		if published, err := post.Publish(); err != nil {
			t.Fatal("failed publish", err)
		} else {
			if published.ID == 0 || published.ID == post.ID {
				t.Fatal("failed publish", published)
			}
			var posts Posts
			if err := DB.Where("post_parent_id = ?", published.PostParentID).Where("type = ?", Published).Preload("Tags").Find(&posts).Error; err != nil {
				t.Fatal("failed publish", err)
			}
			if len(posts) != 1 {
				t.Fatal("failed publish", posts)
			}
			if len(posts[0].Tags) != 1 {
				t.Fatal("failed publish", posts[0])
			}
			if err := DB.Where("post_parent_id = ?", published.PostParentID).Where("type = ?", PublishedLog).Find(&posts).Error; err != nil {
				t.Fatal("failed publish", err)
			}
			if len(posts) != 1 {
				t.Fatal("failed publish", posts)
			}
		}
	})
}

func TestUnpublish(t *testing.T) {
	postListWrapper(1, func(posts Posts) {
		post := posts[0]
		if _, err := post.Publish(); err != nil {
			t.Fatal("failed publish", err)
		}
		if err := post.Unpublish(); err != nil {
			t.Fatal("failed publish", err)
		}
		if err := DB.Where("post_parent_id = ?", post.PostParentID).Where("type = ?", Published).Find(&posts).Error; err != nil {
			t.Fatal("failed publish", err)
		}
		if len(posts) != 0 {
			t.Fatal("failed publish", posts)
		}
		if err := DB.Where("post_parent_id = ?", post.PostParentID).Where("type = ?", PublishedLog).Find(&posts).Error; err != nil {
			t.Fatal("failed publish", err)
		}
		if len(posts) != 1 {
			t.Fatal("failed publish", posts)
		}
	})
}

func TestIndexPost(t *testing.T) {
	postListWrapper(3, func(_ Posts) {
		posts := Posts{}
		posts.Index(nil)
		if len(posts) != 3 {
			t.Fatal("failed test", len(posts))
		}
		if len(posts[0].Tags) != 1 {
			t.Fatal("failed test", posts[0].Tags)
		}
	})
}

func TestIndexPostWithType(t *testing.T) {
	postListWrapper(3, func(_posts Posts) {
		draft := Post{Title: "published", UserID: _posts[0].UserID}
		if err := draft.Create(); err != nil {
			t.Fatal("failed test", err)
		}
		if _, err := draft.Publish(); err != nil {
			t.Fatal("failed test", err)
		}
		posts := Posts{}
		posts.Index(&PostPagination{PostType: Published})
		if len(posts) != 1 {
			t.Fatal("failed test", len(posts))
		}
		if posts[0].Type != Published {
			t.Fatal("failed test", posts[0])
		}
	})
}

func TestIndexPostInUser(t *testing.T) {
	GormOpen()
	user1 := User{Name: "user1", Password: "password"}
	DB.Create(&user1)
	defer DB.Delete(&user1)
	user2 := User{Name: "user2", Password: "password"}
	DB.Create(&user2)
	defer DB.Delete(&user2)
	for i := 0; i < 10; i++ {
		post := Post{
			UserID:   user2.ID,
			Title:    fmt.Sprintf("title_%d", i),
			Problem:  "problem",
			Solution: "solution",
			Lesson:   "lesson",
		}
		if err := post.Create(); err != nil {
			t.Fatal("failed test", err)
		}
		defer DB.Delete(post.PostParent)
	}
	for i := 0; i < 10; i++ {
		post := Post{
			UserID:   user1.ID,
			Title:    fmt.Sprintf("title_%d", i),
			Problem:  "problem",
			Solution: "solution",
			Lesson:   "lesson",
		}
		if err := post.Create(); err != nil {
			t.Fatal("failed test", err)
		}
		defer DB.Delete(post.PostParent)
	}

	posts := Posts{}
	if err := posts.Index(&PostPagination{
		Pagination: Pagination{
			Page:    2,
			Limit:   3,
			OrderBy: "id",
			Keyword: "",
		},
		UserID: user2.ID,
	}); err != nil {
		t.Fatal("failed")
	}
	if len(posts) != 3 {
		t.Fatal("failed", posts)
	}
	if posts[0].UserID != user2.ID {
		t.Fatal("failed", posts)
	}
}

func TestIndexPostWithTag(t *testing.T) {
	GormOpen()
	user1 := User{Name: "user1", Password: "password"}
	DB.Create(&user1)
	defer DB.Delete(&user1)
	for i := 0; i < 5; i++ {
		post := Post{
			UserID:   user1.ID,
			Title:    fmt.Sprintf("title_%d", i),
			Problem:  "problem",
			Solution: "solution",
			Lesson:   "lesson",
		}
		if err := post.Create(); err != nil {
			t.Fatal("failed test", err)
		}
		defer DB.Delete(post.PostParent)
	}
	tag := Tag{Title: "tag"}
	DB.Create(&tag)
	defer DB.Delete(&tag)
	tags := []Tag{tag}
	for i := 5; i < 10; i++ {
		post := Post{
			UserID:   user1.ID,
			Title:    fmt.Sprintf("title_%d", i),
			Problem:  "problem",
			Solution: "solution",
			Lesson:   "lesson",
			Tags:     tags,
		}
		if err := post.Create(); err != nil {
			t.Fatal("failed test", err)
		}
		defer DB.Delete(post.PostParent)
	}

	posts := Posts{}
	if err := posts.Index(&PostPagination{
		Tag: "tag",
	}); err != nil {
		t.Fatal("failed", err)
	}
	if len(posts) != 5 {
		t.Fatal("failed", posts)
	}
	if posts[0].Tags[0].Title != "tag" {
		t.Fatal("failed", posts)
	}
}

func TestBatchDeletePost(t *testing.T) {
	postListWrapper(3, func(created Posts) {
		BatchDeletePost([]int{created[0].ID, created[1].ID, created[2].ID})
		posts := Posts{}
		posts.Index(nil)
		if len(posts) != 0 {
			t.Fatal("failed delete", len(posts))
		}
	})
}
