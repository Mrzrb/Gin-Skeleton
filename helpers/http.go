package helpers

import (
	"app/config"

	"github.com/gin-gonic/gin"
)

func HttpServer(engine *gin.Engine) {
	// 初始化http服务路由
	// router.Http(engine)
	// // MQ 消费回调路由
	// router.MQ(engine)
	// app内定时任务
	// router.Tasks(engine)

	// 启动web server
	engine.Run(config.Conf.Server.Addr)
}
