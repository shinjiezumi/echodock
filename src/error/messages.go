package error

// Message エラ〜メッセージ
type Message string

const (
	MessageBadRequest       Message = "Bad Request"
	MessageResourceNotFound Message = "Resource not found"
)
