package tools

import (
	"app/infra"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ErrNo  int64  `json:"errNo"`
	ErrMsg string `json:"errMsg"`
	Data   any    `json:"data"`
}

type EmptyRsp struct{}

func RenderJsonSucc(ctx *gin.Context, data any) {
	rsp := Response{0, "succ", data}
	renderJsonSucc(ctx, rsp)
}

func renderJsonSucc(ctx *gin.Context, rsp Response) {
	ctx.JSON(http.StatusOK, rsp)
}

func renderJsonFail(ctx *gin.Context, rsp Response) {
	ctx.JSON(http.StatusOK, rsp)
}

func RenderJsonFail(ctx *gin.Context, err error) {
	rsp := Response{400, err.Error(), EmptyRsp{}}
	if serr, ok := err.(*infra.BizError); ok {
		rsp.ErrNo = serr.Code
	}
	renderJsonFail(ctx, rsp)
}
