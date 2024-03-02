package router

import (
	"github.com/gin-gonic/gin"
)

func Http(engine *gin.Engine) {
	engine.Any("/api", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"code": 0,
			"data": "success",
		})
	})
}
