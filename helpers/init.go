package helpers

import (
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	InitMysql()
}
