package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Post ポスト
type Post struct {
	Model
	UserID       int         `json:"userID" binding:"required"`
	User         *User       `json:"user" gorm:"foreignKey:UserID"`
	PostParentID int         `json:"postParentID"`
	PostParent   *PostParent `json:"postParent" gorm:"foreignKey:PostParentID"`
	Title        string      `json:"title" binding:"required,lte=256"`
	Problem      string      `json:"problem"`
	Solution     string      `json:"solution"`
	Lesson       string      `json:"lesson"`
	Type         PostType    `json:"type"` // 1: 下書き 2: 公開 3: 公開履歴
	Tags         Tags        `json:"tags" gorm:"many2many:post_tags;"`
}

// PostType ポスト種別
type PostType int

const (
	// Draft 下書き
	Draft PostType = 1
	// Published 公開
	Published PostType = 2
	// PublishedLog 公開ログ
	PublishedLog PostType = 3
)

// Posts ポスト一覧
type Posts []Post

// PostPagination ポストページネーション条件
type PostPagination struct {
	Pagination
	UserID   int
	Tag      string
	PostType PostType
}

// BeforeSave バリデーション
func (post *Post) BeforeSave() error {
	return GetValidator().Struct(post)
}

// Create 作成
func (post *Post) Create() error {
	post.Type = Draft
	return Tx(func(db *gorm.DB) error {
		postParent := PostParent{Status: PostParentDraft}
		if err := postParent.Create(); err != nil {
			return err
		}
		post.PostParentID = postParent.ID
		post.PostParent = &postParent
		return DB.Create(post).Error
	})
}

// Read 読込
func (post *Post) Read() error {
	if post.ID != 0 {
		postParent := &PostParent{}
		post.PostParent = postParent
		return DB.First(post).Related(&post.Tags, "Tags").Related(post.PostParent).Error
	}
	return errors.New("no key to read")
}

// Update 更新
func (post *Post) Update() error {
	if post.Type != Draft {
		return errors.New("only draft can be updated")
	}
	tags := post.Tags
	return Tx(func(db *gorm.DB) error {
		return db.Save(post).Model(post).Association("Tags").Replace(tags).Error
	})
}

// Delete 削除
func (post *Post) Delete() error {
	if post.Type == Draft {
		return DB.Delete(post.PostParent).Error
	}
	return DB.Delete(post).Error
}

// Publish 公開
func (post *Post) Publish() (*Post, error) {
	if post.Type != Draft {
		return nil, errors.New("only draft can be published")
	}
	published := Post{
		UserID:       post.UserID,
		Title:        post.Title,
		Problem:      post.Problem,
		Solution:     post.Solution,
		Lesson:       post.Lesson,
		Tags:         post.Tags,
		Type:         Published,
		PostParentID: post.PostParentID,
	}
	if err := Tx(func(db *gorm.DB) error {
		if err := db.Table("posts").Where("post_parent_id = ?", post.PostParentID).Where("type = ?", Published).Updates(map[string]interface{}{"type": PublishedLog}).Error; err != nil {
			return err
		}
		if err := db.Create(&published).Error; err != nil {
			return err
		}
		post.PostParent.Status = PostParentPublished
		return db.Save(post.PostParent).Error
	}); err != nil {
		return nil, err
	}
	return &published, nil
}

// Unpublish 公開中止
func (post *Post) Unpublish() error {
	return Tx(func(db *gorm.DB) error {
		if err := db.Table("posts").Where("post_parent_id = ?", post.PostParentID).Where("type = ?", Published).Updates(map[string]interface{}{"type": PublishedLog}).Error; err != nil {
			return err
		}
		post.PostParent.Status = PostParentDraft
		return db.Save(post.PostParent).Error
	})
}

func filterType(db *gorm.DB, postType PostType) *gorm.DB {
	if postType == Draft {
		// 下書きしかポストが存在しない
		return db.Joins("INNER JOIN post_parents ON post_parents.id = posts.post_parent_id AND post_parents.status = 1 AND posts.type = 1")
	} else if postType == Published {
		// 公開ポストが存在する
		return db.Joins("INNER JOIN post_parents ON post_parents.id = posts.post_parent_id AND post_parents.status = 2 AND posts.type = 2")
	}
	return db
}

// Index 一覧
func (posts *Posts) Index(pagination *PostPagination) error {
	db := DB
	if pagination != nil {
		db = paginate(db, &pagination.Pagination)
		db = filterType(db, pagination.PostType)
		if pagination.UserID != 0 {
			db = db.Where("user_id = ?", pagination.UserID)
		}
		if pagination.Tag != "" {
			db = db.Joins("INNER JOIN post_tags ON post_tags.post_id = posts.id")
			db = db.Joins("INNER JOIN tags ON post_tags.tag_id = tags.id AND tags.title = ?", pagination.Tag)
		}
	}
	db = db.Preload("Tags")
	db = db.Preload("PostParent")
	return db.Find(posts).Error
}

// BatchDeletePost 一覧削除
func BatchDeletePost(ids []int) error {
	return DB.Where(ids).Delete(Post{}).Error
}
