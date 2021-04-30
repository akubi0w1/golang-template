package response

type User struct {
	ID        int     `json:"id"`
	AccountID string  `json:"accountId"`
	Email     string  `json:"email"`
	Profile   Profile `json:"profile"`
}

type Profile struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
