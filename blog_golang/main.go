package main

import (
	"blog/config"
	"blog/routes"
)

func main() {
	// 打开数据库连接
	config.OpenDB()

	r := routes.SetupRoutes()
	r.Run(":9999")

	// 起初测试创建Gin引擎
	// r := gin.Default()

	// 健康检查
	//r.GET("/health", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Healthy",
	//	})
	//})

	//r.Run()
}
