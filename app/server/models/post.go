package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Post ポスト
type Post struct {
	Model
	UserID   int    `json:"userID" binding:"required"`
	User     *User  `json:"user" gorm:"ForeignKey:UserID"`
	Title    string `json:"title" binding:"required,lte=256"`
	Problem  string `json:"problem"`
	Solution string `json:"solution"`
	Lesson   string `json:"lesson"`
	Tags     Tags   `json:"tags" gorm:"many2many:post_tags;"`
}

// Posts ポスト一覧
type Posts []Post

// PostPagination ポストページネーション条件
type PostPagination struct {
	Pagination
	UserID int
	Tag    string
}

// BeforeSave バリデーション
func (post *Post) BeforeSave() error {
	return GetValidator().Struct(post)
}

// Create 作成
func (post *Post) Create() error {
	return DB.Create(post).Error
}

// Read 読込
func (post *Post) Read() error {
	if post.ID != 0 {
		return DB.Preload("Tags").First(post).Error
	}
	return errors.New("no key to read")
}

// Update 更新
func (post *Post) Update() error {
	tags := post.Tags
	return Tx(func(db *gorm.DB) error {
		return db.Save(post).Model(post).Association("Tags").Replace(tags).Error
	})
}

// Delete 削除
func (post *Post) Delete() error {
	return DB.Delete(post).Error
}

// Index 一覧
func (posts *Posts) Index(pagination *PostPagination) error {
	db := DB
	if pagination != nil {
		db = paginate(db, &pagination.Pagination)
		if pagination.UserID != 0 {
			db = db.Where("user_id = ?", pagination.UserID)
		}
		if pagination.Tag != "" {
			db = db.Joins("INNER JOIN post_tags ON post_tags.post_id = posts.id")
			db = db.Joins("INNER JOIN tags ON post_tags.tag_id = tags.id AND tags.title = ?", pagination.Tag)
		}
		db = db.Preload("Tags")
	} else {
		db = db.Preload("Tags")
	}
	return db.Find(posts).Error
}

// BatchDeletePost 一覧削除
func BatchDeletePost(ids []int) error {
	return DB.Where(ids).Delete(Post{}).Error
}
