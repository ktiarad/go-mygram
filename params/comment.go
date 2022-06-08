package params

type CommentCreate struct {
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
	UserID  int    `json:"user_id"`
}

type CommentUpdate struct {
	Message string `json:"message"`
}
