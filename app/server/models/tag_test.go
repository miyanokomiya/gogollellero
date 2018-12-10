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
	if tag.Title != "title" {
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
	tag1 := Tag{
		Title: "title1",
	}
	tag2 := Tag{
		Title: "title2",
	}
	var tags Tags
	tags = append(tags, tag1, tag2)
	defer DB.Delete(&tag1)
	defer DB.Delete(&tag2)
	if err := tags.CreateIfNotExist(); err != nil {
		t.Fatal("failed test", err)
	}
	// 再作成
	if err := tags.CreateIfNotExist(); err != nil {
		t.Fatal("failed test", err)
	}
}
