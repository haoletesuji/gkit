package pkg

import (
	gkit "gkit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *HttpServer) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (s *HttpServer) Error(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gkit.ErrResponse{
		Success: false,
		Error:   "response error",
	})
}

func (s *HttpServer) Success(ctx *gin.Context) {
	users := make([]User, 0)

	users = append(users, User{
		UID:          "1",
		Name:         "John",
		ProfileImage: "https://google.com/image/1",
	})

	ctx.JSON(http.StatusOK, &gkit.SuccessResponse[[]User]{
		Success: true,
		Data:    users,
	})
}

func (s *HttpServer) SuccessPaging(ctx *gin.Context) {
	users := make([]User, 0)

	users = append(users, User{
		UID:          "1",
		Name:         "John",
		ProfileImage: "https://google.com/image/1",
	})

	ctx.JSON(http.StatusOK, &gkit.SuccessPagingResponse[[]User]{
		Success: true,
		Data:    users,
		Pagination: gkit.Pagination{
			Total: 1,
			Limit: 10,
		},
	})
}
