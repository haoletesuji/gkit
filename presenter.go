package common

type ErrResponse struct {
	Success bool   `json:"success" example:"true"`
	Error   string `json:"error"`
}

type SuccessResponse[T any] struct {
	Success bool `json:"success" example:"true"`
	Data    T    `json:"data"`
}

type Pagination struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
}

type SuccessPagingResponse[T any] struct {
	Success    bool       `json:"success" example:"true"`
	Data       T          `json:"data"`
	Pagination Pagination `json:"pagination"`
}
