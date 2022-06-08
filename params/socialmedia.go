package params

type SocialMediaCreate struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int    `json:"user_id"`
}
