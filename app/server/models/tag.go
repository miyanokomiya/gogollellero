package models

import (
	"github.com/jinzhu/gorm"
)

// Tag タグ
type Tag struct {
	Model
	Title string `json:"title" binding:"required,lte=256"`
}

// Tags タグ一覧
type Tags []Tag

// CreateIfNotExist 未作成ならば作成
func (tag *Tag) CreateIfNotExist() error {
	return DB.Where(Tag{Title: tag.Title}).FirstOrCreate(tag).Error
}

// CreateTagsIfNotExist 未作成ならば作成
func CreateTagsIfNotExist(titles []string) (Tags, error) {
	var tags Tags
	var tag Tag
	for _, title := range titles {
		tag = Tag{Title: title}
		if err := tag.CreateIfNotExist(); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func tagsWhere(db *gorm.DB, pagination *Pagination) *gorm.DB {
	if pagination != nil && pagination.Keyword != "" {
		db = DB.Where("title LIKE ?", "%"+pagination.Keyword+"%")
	}
	return db
}

// Index 一覧
func (tags *Tags) Index(pagination *Pagination) error {
	return paginate(tagsWhere(DB, pagination), pagination).Find(tags).Error
}

// IndexOfUser 一覧 ユーザー指定
func (tags *Tags) IndexOfUser(userID int) error {
	return DB.Raw(`
SELECT tags.*
FROM posts
  INNER JOIN post_tags
    ON posts.id = post_tags.post_id
    AND posts.user_id = ? 
  INNER JOIN tags ON post_tags.tag_id = tags.id
GROUP BY tags.id
	`, userID).Scan(&tags).Error
}
