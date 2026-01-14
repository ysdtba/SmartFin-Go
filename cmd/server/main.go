package main

import (
	"log"

	"github.com/florentyang/smartfin-go/internal/bootstrap"
	"github.com/florentyang/smartfin-go/internal/router"
)

func main() {
	// 1. 初始化应用（所有依赖注入在 bootstrap 里完成）
	app := bootstrap.NewApp()

	// 2. 设置路由（传入 Controller）
	r := router.SetupRouter(app.UserController)

	// 3. 启动服务器
	log.Println("====================================")
	log.Println("🚀 SmartFin-Go 服务启动中...")
	log.Println("📍 访问地址: http://localhost:8080")
	log.Println("====================================")
	log.Println("📋 API 列表:")
	log.Println("   GET  /health              - 健康检查")
	log.Println("   POST /api/v1/user/register - 用户注册")
	log.Println("   POST /api/v1/user/login    - 用户登录")
	log.Println("====================================")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
