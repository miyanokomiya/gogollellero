package models

// PostParent ポスト親データ
type PostParent struct {
	Model
	ViewCount int `json:"viewCount"`
}

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
