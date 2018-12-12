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

func tagsWhere(db *gorm.DB, keyword string) *gorm.DB {
	if keyword != "" {
		db = DB.Where("title LIKE ?", "%"+keyword+"%")
	}
	return db
}

// Index 一覧
func (tags *Tags) Index(pagination *Pagination) error {
	return paginate(tagsWhere(DB, pagination.Keyword), pagination).Find(tags).Error
}

// IndexOfUser 一覧 ユーザー指定
func IndexOfUser(userID int) (Tags, error) {
	var tags Tags
	err := DB.Raw("SELECT tags.* FROM post_tags LEFT JOIN tags ON post_tags.tag_id = tags.id LEFT JOIN posts ON posts.id = post_tags.post_id WHERE posts.user_id = ? GROUP BY tags.title", userID).Scan(&tags).Error
	return tags, err
}
