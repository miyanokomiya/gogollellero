package models

import (
	"testing"
)

func TestCreatePostParent(t *testing.T) {
	GormOpen()
	postParent := PostParent{}
	if err := postParent.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer DB.Delete(&postParent)

	if postParent.ID == 0 {
		t.Fatal("failed test", postParent)
	}
}

func TestDeletePostParent(t *testing.T) {
	GormOpen()
	postParent1 := PostParent{}
	if err := postParent1.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer DB.Delete(&postParent1)
	postParent2 := PostParent{}
	if err := postParent2.Create(); err != nil {
		t.Fatal("failed test", err)
	}
	defer DB.Delete(&postParent2)
	if err := postParent2.Delete(); err != nil {
		t.Fatal("failed test", err)
	}
	deleted := PostParent{}
	DB.First(&deleted, postParent2.ID)
	if deleted.ID != 0 {
		t.Fatal("failed delete", deleted)
	}
}
