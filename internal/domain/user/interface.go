package user

import (
	"errors"

	"github.com/florentyang/smartfin-go/internal/entity"
)

// ==================== 错误定义 ====================
// 领域层的业务错误，供上层判断使用

var (
	ErrUserNotFound     = errors.New("用户不存在")
	ErrUsernameExists   = errors.New("用户名已存在")
	ErrEmailExists      = errors.New("邮箱已存在")
	ErrInvalidPassword  = errors.New("密码错误")
	ErrPasswordTooShort = errors.New("密码至少6个字符")
)

// ==================== Domain 接口定义 ====================
// Service 层会依赖这个接口，而不是具体实现

type Domain interface {
	// Register 注册新用户
	Register(username, email, password string) (*entity.User, error)

	// Login 用户登录
	Login(username, password string) (*entity.User, error)

	// GetProfile 获取用户个人信息
	GetProfile(userID uint) (*entity.User, error)

	// UpdateProfile 更新用户个人信息
	UpdateProfile(userID uint, username, email string) error

	// UpdatePassword 更新用户密码（需验证旧密码）
	UpdatePassword(userID uint, oldPassword, newPassword string) error
}
