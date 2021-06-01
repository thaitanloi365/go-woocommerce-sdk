package woocommerce

type ApiError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func (e *ApiError) Error() string {
	return e.Message
}
