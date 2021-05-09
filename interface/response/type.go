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

type UserList struct {
	Total int    `json:"total"`
	Users []User `json:"users"`
}

type Token struct {
	Token string `json:"token"`
}

type Post struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	AuthorID  int     `json:"authorId"`
	Tags      []Tag   `json:"tags"`
	Images    []Image `json:"images"`
}

type PostList struct {
	Total int    `json:"total"`
	Posts []Post `json:"posts"`
}

type Tag struct {
	ID  int    `json:"id"`
	Tag string `json:"tag"`
}

type Image struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
