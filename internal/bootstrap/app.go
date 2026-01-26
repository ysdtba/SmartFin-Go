package bootstrap

import (
	"log"

	"gorm.io/gorm"

	"github.com/florentyang/smartfin-go/internal/config"
	"github.com/florentyang/smartfin-go/internal/controller"
	txRepoImpl "github.com/florentyang/smartfin-go/internal/dao/transaction/impl"
	userRepoImpl "github.com/florentyang/smartfin-go/internal/dao/user/impl"
	txDomainImpl "github.com/florentyang/smartfin-go/internal/domain/transaction/impl"
	userDomainImpl "github.com/florentyang/smartfin-go/internal/domain/user/impl"
	"github.com/florentyang/smartfin-go/internal/service"
)

// App 应用程序结构体，包含所有依赖
type App struct {
	DB *gorm.DB

	// Controllers（给 Router 用）
	UserController        controller.UserController
	TransactionController controller.TransactionController
}

// NewApp 创建并初始化应用程序
// 所有依赖注入都在这里完成
func NewApp() *App {
	app := &App{}

	// ==================== 1. 基础设施层 ====================
	app.initDatabase()

	// ==================== 2. 业务层初始化 ====================
	app.initUserModule()

	app.initTransactionModule()
	// TODO: 以后加其他模块
	// app.initAssetModule()
	// app.initTransactionModule()

	return app
}

// initDatabase 初始化数据库连接
func (app *App) initDatabase() {
	db, err := config.InitDB(config.DefaultConfig())
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	app.DB = db
}

// initUserModule 初始化用户模块（依赖注入链）
func (app *App) initUserModule() {
	// DAO → Domain → Service → Controller
	userRepo := userRepoImpl.NewUserRepo(app.DB)         // 使用 DAO impl 包
	userDomain := userDomainImpl.NewUserDomain(userRepo) // 使用 Domain impl 包
	userService := service.NewUserService(userDomain)
	userController := controller.NewUserController(userService)

	app.UserController = userController
}

func (app *App) initTransactionModule() {
	txRepo := txRepoImpl.NewTransactionRepo(app.DB)
	txDomain := txDomainImpl.NewTransactionDomain(txRepo)
	txService := service.NewTransactionService(txDomain)
	txController := controller.NewTransactionController(txService)

	app.TransactionController = txController
}
