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

func TestGetChildDraft(t *testing.T) {
	postListWrapper(3, func(posts Posts) {
		post, err := posts[0].PostParent.GetChild(Draft)
		if err != nil {
			t.Fatal("failed test", err)
		}
		if post == nil {
			t.Fatal("failed test")
		}
		if post.PostParent == nil {
			t.Fatal("failed test", post)
		}
	})
}

func TestGetChildPublished(t *testing.T) {
	postListWrapper(3, func(posts Posts) {
		if _, err := posts[0].Publish(); err != nil {
			t.Fatal("failed test", err)
		}
		post, err := posts[0].PostParent.GetChild(Published)
		if err != nil {
			t.Fatal("failed test", err)
		}
		if post == nil {
			t.Fatal("failed test")
		}
		if post.PostParent == nil {
			t.Fatal("failed test", post)
		}
	})
}
