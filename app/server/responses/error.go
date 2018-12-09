package responses

// Error エラーレスポンス
type Error struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}
