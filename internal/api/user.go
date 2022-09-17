package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kjhch/gin-bench/internal/service"
	"net/http"
	"strconv"
)

type Response[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type UserApi struct {
	userService *service.UserService
}

func NewUserApi(userService *service.UserService) *UserApi {
	return &UserApi{userService: userService}
}

func (api *UserApi) GetUser(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	ctx.JSON(http.StatusOK, successfulResp(api.userService.GetUser(id)))
}

func (api *UserApi) ListUsers(ctx *gin.Context) {
	pageNum, _ := strconv.ParseUint(ctx.Query("pageNum"), 10, 32)
	pageSize, _ := strconv.ParseUint(ctx.Query("pageSize"), 10, 32)
	ctx.JSON(http.StatusOK, successfulResp(api.userService.ListUsers(uint(pageNum), uint(pageSize))))

}
func successfulResp[T any](data T) *Response[T] {
	return &Response[T]{
		Code:    "00000",
		Message: "success",
		Data:    data,
	}
}
