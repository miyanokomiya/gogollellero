package models

// PostParent ポスト親データ
type PostParent struct {
	Model
	Status    PostParentStatus `json:"status"`
	ViewCount int              `json:"viewCount"`
}

// PostParentStatus ポスト親データ状態
type PostParentStatus int

const (
	// PostParentDraft 下書き
	PostParentDraft PostParentStatus = 1
	// PostParentPublished 公開
	PostParentPublished PostParentStatus = 2
)

// PostParents ポスト親データ一覧
type PostParents []PostParent

// Create 作成
func (postParent *PostParent) Create() error {
	return DB.Create(postParent).Error
}

// Delete 削除
func (postParent *PostParent) Delete() error {
	return DB.Delete(postParent).Error
}

// GetChild 子供ポスト取得
func (postParent *PostParent) GetChild(postType PostType) (post *Post, err error) {
	post = &Post{}
	if err = DB.Where("post_parent_id = ? AND type = ?", postParent.ID, postType).First(post).Error; err != nil {
		return nil, err
	}
	if err = post.Read(); err != nil {
		return nil, err
	}
	return post, err
}
