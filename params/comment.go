package params

type CommentCreate struct {
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
}

type CommentUpdate struct {
	Message string `json:"message"`
}
