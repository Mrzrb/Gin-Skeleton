package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	errNo  int64
	errMsg string
	data   any
}

type EmptyRsp struct{}

func RenderJsonSucc(ctx *gin.Context, data any) {
	rsp := Response{0, "succ", data}
	ctx.JSON(http.StatusOK, rsp)
}

func RenderJsonFail(ctx *gin.Context, err error) {
	rsp := Response{-1, err.Error(), EmptyRsp{}}
	ctx.JSON(http.StatusOK, rsp)
}
