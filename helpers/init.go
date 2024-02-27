package helpers

import (
	"app/router"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	router.Http(engine)
	InitMysql()
}
