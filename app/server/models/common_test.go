package models

import (
	"testing"
)

func TestGormConnect(t *testing.T) {
	db := gormConnect()
	if db == nil {
		t.Fatal("failed test")
	}
}
