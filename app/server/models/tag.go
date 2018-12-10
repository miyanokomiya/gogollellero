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
