package models

// Tag タグ
type Tag struct {
	Model
	Title string `json:"title" binding:"required,lte=256"`
}
