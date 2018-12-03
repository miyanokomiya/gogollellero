package models

import (
	"testing"
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
