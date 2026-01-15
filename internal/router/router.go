package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/florentyang/smartfin-go/internal/controller"
	"github.com/florentyang/smartfin-go/internal/middleware"
)

// SetupRouter 初始化并配置所有路由
// 参数：userController 从 bootstrap 传入
func SetupRouter(userController controller.UserController) *gin.Engine {
	r := gin.Default()

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "SmartFin-Go is running! 🚀",
		})
	})

	// ==================== 公开接口（无需登录） ====================
	publicGroup := r.Group("/api/v1/user")
	{
		publicGroup.POST("/register", userController.Register)
		publicGroup.POST("/login", userController.Login)
	}

	// ==================== 私有接口（需要 JWT 鉴权） ====================
	authGroup := r.Group("/api/v1/user")
	authGroup.Use(middleware.JWTAuth()) // ← 使用 JWT 中间件
	{
		authGroup.GET("/profile", userController.GetProfile)       // 获取个人信息
		authGroup.PUT("/profile", userController.UpdateProfile)    // 更新个人信息
		authGroup.POST("/password", userController.UpdatePassword) // 更新密码

	}

	return r
}
