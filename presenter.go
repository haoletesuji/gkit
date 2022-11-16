package gkit

var (
	DefaultLimit int64 = 20
)

type ErrResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error"`
}

type SuccessResponse[T any] struct {
	Success bool `json:"success" example:"true"`
	Data    T    `json:"data"`
}

type Pagination struct {
	Total int64 `json:"total"`
	Limit int64 `json:"limit"`
}

type SuccessPagingResponse[T any] struct {
	Success    bool       `json:"success" example:"true"`
	Data       T          `json:"data"`
	Pagination Pagination `json:"pagination"`
}
