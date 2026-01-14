package service

import (
	userDomain "github.com/florentyang/smartfin-go/internal/domain/user"
	"github.com/florentyang/smartfin-go/internal/dto"
	"github.com/florentyang/smartfin-go/internal/entity"
	"github.com/florentyang/smartfin-go/pkg/jwt"
)

// ==================== 接口定义 ====================
// Controller 层会使用这个接口

type UserService interface {
	Register(req *dto.RegisterRequest) (*dto.UserResponse, error)
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	GetProfile(userID uint) (*dto.UserResponse, error)
	UpdateProfile(userID uint, req *dto.UpdateUserRequest) error
}

// ==================== 接口实现 ====================

type userService struct {
	userDomain userDomain.Domain // import Domain 层的接口
}

// NewUserService 创建 Service 实例
func NewUserService(userDomain userDomain.Domain) UserService {
	return &userService{
		userDomain: userDomain,
	}
}

// Register 用户注册
// Service 层职责：协调调用 + DTO 转换
func (s *userService) Register(req *dto.RegisterRequest) (*dto.UserResponse, error) {
	// 1. 调用 Domain 层处理核心业务逻辑
	user, err := s.userDomain.Register(req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	// 2. Entity → DTO 转换（不暴露内部结构给前端）
	return entityToDTO(user), nil
}

// ==================== 私有辅助函数 ====================

// entityToDTO 将 Entity 转换为 DTO（隐藏敏感字段如密码）
func entityToDTO(user *entity.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// Login 用户登录
// Service 层职责：调用 Domain 验证 + 生成 JWT Token
func (s *userService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1. 调用 Domain 层验证用户名和密码
	user, err := s.userDomain.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// 2. 生成 JWT Token（Service 层的职责）
	token, expiresAt, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	// 3. 组装登录响应
	return &dto.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      *entityToDTO(user),
	}, nil
}

// GetProfile 获取用户个人信息
// Service 层职责：调用 Domain 层获取用户信息
func (s *userService) GetProfile(userID uint) (*dto.UserResponse, error) {
	// 1. 调用 Domain 层获取用户信息
	user, err := s.userDomain.GetProfile(userID)
	if err != nil {
		return nil, err
	}
	return entityToDTO(user), nil
}

// UpdateProfile 更新用户个人信息
// Service 层职责：调用 Domain 层更新用户信息
func (s *userService) UpdateProfile(userID uint, req *dto.UpdateUserRequest) error {
	// 1. 调用 Domain 层更新用户信息
	err := s.userDomain.UpdateProfile(userID, req.Username, req.Email)
	if err != nil {
		return err
	}
	return nil
}
