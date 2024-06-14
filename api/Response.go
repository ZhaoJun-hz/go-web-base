package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Response struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  int64  `json:"total,omitempty"`
}

func (m Response) IsEmpty() bool {
	return reflect.DeepEqual(m, Response{})
}

func buildStatus(resp Response, defaultStatus int) int {
	if resp.Status == 0 {
		return defaultStatus
	}
	return resp.Status
}

func OK(ctx *gin.Context, resp Response) {
	HttpResponse(ctx, buildStatus(resp, http.StatusOK), resp)
}

func Fail(ctx *gin.Context, resp Response) {
	HttpResponse(ctx, buildStatus(resp, http.StatusBadRequest), resp)
}

func ServerFail(ctx *gin.Context, resp Response) {
	HttpResponse(ctx, buildStatus(resp, http.StatusInternalServerError), resp)
}

func HttpResponse(ctx *gin.Context, status int, resp Response) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	ctx.AbortWithStatusJSON(status, resp)
}
