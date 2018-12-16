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
