package dto

import (
	"time"
)

// ================== 请求 DTO ==================

// 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 更新用户信息请求（基础）
type UpdateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// 更新用户密码请求
type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

// ================== 响应 DTO ==================

// 用户响应
type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// 登录响应（包含 Token）
type LoginResponse struct {
	Token     string       `json:"token"`      // JWT Token
	ExpiresAt int64        `json:"expires_at"` // 过期时间戳
	User      UserResponse `json:"user"`       // 用户信息
}
