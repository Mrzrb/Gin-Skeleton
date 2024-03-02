package infra

import "github.com/gin-gonic/gin"

type Controller interface {
	Routes(engine *gin.Engine)
}
