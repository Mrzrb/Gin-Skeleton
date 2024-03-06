package infra

import (
	"github.com/gin-gonic/gin"
)

type ControllerMethod func(ctx *gin.Context) (any, error)

type Controller interface {
	Routes(engine *gin.Engine)
}
