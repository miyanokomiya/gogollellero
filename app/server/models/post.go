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
func (posts *Posts) Index(pagination *Pagination) error {
	return paginate(DB, pagination).Preload("Tags").Find(posts).Error
}

// IndexInUser 一覧
func (posts *Posts) IndexInUser(pagination *Pagination, userID int) error {
	return paginate(DB.Where("user_id = ?", userID), pagination).Preload("Tags").Find(posts).Error
}

// BatchDeletePost 一覧削除
func BatchDeletePost(ids []int) error {
	return DB.Where(ids).Delete(Post{}).Error
}
