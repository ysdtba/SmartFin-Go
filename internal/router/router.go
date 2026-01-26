package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/florentyang/smartfin-go/internal/controller"
	"github.com/florentyang/smartfin-go/internal/middleware"
)

// SetupRouter 初始化并配置所有路由
// 参数：从 bootstrap 传入各个 Controller
func SetupRouter(
	userController controller.UserController,
	txController controller.TransactionController,
) *gin.Engine {
	r := gin.Default()

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "SmartFin-Go is running! 🚀",
		})
	})

	// ==================== 用户模块 - 公开接口 ====================
	publicGroup := r.Group("/api/v1/user")
	{
		publicGroup.POST("/register", userController.Register)
		publicGroup.POST("/login", userController.Login)
	}

	// ==================== 用户模块 - 私有接口 ====================
	userAuthGroup := r.Group("/api/v1/user")
	userAuthGroup.Use(middleware.JWTAuth())
	{
		userAuthGroup.GET("/profile", userController.GetProfile)       // 获取个人信息
		userAuthGroup.PUT("/profile", userController.UpdateProfile)    // 更新个人信息
		userAuthGroup.POST("/password", userController.UpdatePassword) // 更新密码
	}

	// ==================== 交易模块 - 私有接口 ====================
	txGroup := r.Group("/api/v1/transactions")
	txGroup.Use(middleware.JWTAuth())
	{
		txGroup.POST("/create", txController.Create) // 创建交易：POST /api/v1/transactions/create
	}

	return r
}
