package models

import (
	"testing"
)

func TestCreateIfNotExistTag(t *testing.T) {
	GormOpen()
	tag := Tag{
		Title: "title",
	}
	if err := tag.CreateIfNotExist(); err != nil {
		t.Fatal("failed test", err)
	}
	defer DB.Delete(&tag)
	if tag.ID == 0 {
		t.Fatal("failed test", tag)
	}
	// 存在するものを再度保存しても問題なし
	tag2 := Tag{
		Title: "title",
	}
	if err := tag2.CreateIfNotExist(); err != nil {
		t.Fatal("failed test", err)
	}
}

func TestCreateIfNotExistTags(t *testing.T) {
	GormOpen()
	titles := []string{"a", "b"}
	tags, err := CreateTagsIfNotExist(titles)
	if err != nil {
		t.Fatal("failed test", err)
	}
	for _, tag := range tags {
		defer DB.Delete(&tag)
		if tag.ID == 0 {
			t.Fatal("failed test", tag)
		}
	}
	if tags[0].Title != "a" {
		t.Fatal("failed update", tags)
	}
	if tags[1].Title != "b" {
		t.Fatal("failed update", tags)
	}
}

func TestIndexTag(t *testing.T) {
	tagListWrapper(12, func(_ Tags) {
		tags := Tags{}
		tags.Index(&Pagination{
			Keyword: "1",
		})
		if len(tags) != 3 {
			t.Fatal("failed test", tags)
		}
	})
}

func TestIndexTagOfUser(t *testing.T) {
	postListWrapper(3, func(posts Posts) {
		tags, err := IndexOfUser(posts[0].UserID)
		if err != nil {
			t.Fatal("failed test", err)
		}
		if len(tags) != 1 {
			t.Fatal("failed test", tags)
		}
	})
}
