package models

// Tag タグ
type Tag struct {
	Model
	Title string `json:"title" binding:"required,lte=256"`
}

// Tags タグ一覧
type Tags []Tag

// CreateIfNotExist 未作成ならば作成
func (tag *Tag) CreateIfNotExist() error {
	return DB.FirstOrCreate(tag).Error
}

// CreateIfNotExist 未作成ならば作成
func (tags Tags) CreateIfNotExist() error {
	for _, tag := range tags {
		if err := tag.CreateIfNotExist(); err != nil {
			return err
		}
	}
	return nil
}
