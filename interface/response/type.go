package response

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
