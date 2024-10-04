package response

type SuccessResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string            `json:"status"`
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}