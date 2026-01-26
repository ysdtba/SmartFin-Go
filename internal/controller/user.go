package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/florentyang/smartfin-go/internal/dto"
	"github.com/florentyang/smartfin-go/internal/service"
	"github.com/florentyang/smartfin-go/pkg/response"
)

// ==================== 接口定义 ====================
// Router 层会使用这个接口

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	UpdatePassword(c *gin.Context)
}


// ==================== 接口实现 ====================

type userController struct {
	userService service.UserService // import Service 层的接口
}

// NewUserController 创建 Controller 实例
func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// Register 用户注册接口
// POST /api/v1/user/register
func (ctrl *userController) Register(c *gin.Context) {
	// 1. 绑定请求参数
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 2. 调用 Service 层
	user, err := ctrl.userService.Register(&req)
	if err != nil {
		// 根据错误类型返回不同响应
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	// 3. 返回成功响应
	response.Success(c, user)

}

// Login 用户登录接口
// POST /api/v1/user/login
func (ctrl *userController) Login(c *gin.Context) {
	// 1. 绑定请求参数
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 2. 调用 Service 层（验证用户 + 生成 Token）
	loginResp, err := ctrl.userService.Login(&req)
	if err != nil {
		// 根据错误类型返回不同响应
		response.Fail(c, http.StatusUnauthorized, err.Error())
		return
	}

	// 3. 返回成功响应（包含 Token）
	response.Success(c, loginResp)
}

// GetProfile 获取用户个人信息接口
// GET /api/v1/user/profile
// 需要 JWT 鉴权，userID 从中间件获取
func (ctrl *userController) GetProfile(c *gin.Context) {
	// 1. 从 Context 获取 userID（JWT 中间件已验证并存入）
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	// 2. 调用 Service 层获取用户信息
	userResp, err := ctrl.userService.GetProfile(userID.(uint))
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 3. 返回成功响应
	response.Success(c, userResp)
}

// UpdateProfile 更新用户个人信息接口
// PUT /api/v1/user/update-profile
// 需要 JWT 鉴权，userID 从中间件获取
func (ctrl *userController) UpdateProfile(c *gin.Context) {
	// 1. 从 Context 获取 userID
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	// 2. 绑定请求参数 ← 添加这一步！
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 3. 调用 Service 层更新用户信息
	err := ctrl.userService.UpdateProfile(userID.(uint), &req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 4. 返回成功响应
	response.Success(c, "更新成功")
}

// UpdatePassword 更新用户密码接口
// POST /api/v1/user/password
// 需要 JWT 鉴权，userID 从中间件获取
func (ctrl *userController) UpdatePassword(c *gin.Context) {
	// 1. 从 Context 获取 userID
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	// 2. 绑定请求参数
	var req dto.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 3. 调用 Service 层更新密码
	err := ctrl.userService.UpdatePassword(userID.(uint), &req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 4. 返回成功响应
	response.Success(c, "更新成功")
}	


