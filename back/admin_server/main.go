package main

import (
	"fmt"
	"goods/common/util/redis"
	"log"

	"goods/common/logger"
	"goods/common/model"
	"goods/config"
	"goods/web/middleware"
	"goods/web/route"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化配置文件
	config.InitConfig("./config/config.json")
	// 初始化logger
	logger.InitLogger()
	// 初始化mysql
	model.InitMysql()
	// 初始化redis
	redis.InitRedis(config.GlobalConfig.AppConfig)
}

func main() {
	// 获取应用配置文件
	c := config.GlobalConfig.AppConfig
	r := gin.Default()
	// 设置运行模式
	gin.SetMode(c.GinMode)
	// 设置信任代理
	if err := r.SetTrustedProxies(c.TrustedProxies); err != nil {
		log.Fatalln(err)
	}
	r.Use(middleware.Cors())
	// 注册路由
	route.RegisterRoute(r)
	if err := r.Run(fmt.Sprintf(":%d", c.Port)); err != nil {
		log.Fatalln(err)
	}
}
