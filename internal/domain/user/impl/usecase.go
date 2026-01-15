package impl

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	userRepo "github.com/florentyang/smartfin-go/internal/dao/user"
	userDomain "github.com/florentyang/smartfin-go/internal/domain/user"
	"github.com/florentyang/smartfin-go/internal/entity"
)

// ==================== UseCase 结构体（依赖聚合） ====================
// 聚合所有这个模块需要的依赖

type usecase struct {
	userRepo userRepo.Repo // DAO 层接口
	// 以后可以加更多依赖：
	// logger      *logger.Logger
	// redisClient *redis.Client
	// emailClient *email.Client
}

// ==================== 构造函数 ====================

// NewUserDomain 创建 Domain 实例
// 返回接口类型，隐藏实现细节
func NewUserDomain(repo userRepo.Repo) userDomain.Domain {
	return &usecase{
		userRepo: repo,
	}
}

// ==================== 业务方法实现 ====================

// Register 注册新用户（核心业务逻辑）
func (u *usecase) Register(username, email, password string) (*entity.User, error) {
	// 1. 业务规则：密码长度校验
	if len(password) < 6 {
		return nil, userDomain.ErrPasswordTooShort
	}

	// 2. 业务规则：检查用户名是否已存在
	if u.userRepo.ExistsByUsername(username) {
		return nil, userDomain.ErrUsernameExists
	}

	// 3. 业务规则：检查邮箱是否已存在
	if u.userRepo.ExistsByEmail(email) {
		return nil, userDomain.ErrEmailExists
	}

	// 4. 核心逻辑：密码加密
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	// 5. 创建用户对象
	user := &entity.User{
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 6. 调用 DAO 层存储
	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login 用户登录（核心业务逻辑）
func (u *usecase) Login(username, password string) (*entity.User, error) {
	// 1. 根据用户名查找用户
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return nil, userDomain.ErrUserNotFound
	}

	// 2. 验证密码是否正确
	if !checkPassword(user, password) {
		return nil, userDomain.ErrInvalidPassword
	}

	// 3. 登录成功，返回用户信息
	return user, nil
}

// GetProfile 获取用户个人信息
func (u *usecase) GetProfile(userID uint) (*entity.User, error) {
	// 1. 根据用户ID查找用户
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, userDomain.ErrUserNotFound
	}
	return user, nil
}

// UpdateProfile 更新用户个人信息
func (u *usecase) UpdateProfile(userID uint, username, email string) error {
	// 1. 根据用户ID查找用户
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return userDomain.ErrUserNotFound
	}

	// 2. 更新用户信息
	user.Username = username
	user.Email = email
	user.UpdatedAt = time.Now()

	// 3. 调用 DAO 层更新用户信息
	return u.userRepo.Update(user)
}

// UpdatePassword 更新用户密码（验证旧密码 + 加密新密码）
func (u *usecase) UpdatePassword(userID uint, oldPassword, newPassword string) error {
	// 1. 根据用户ID查找用户
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return userDomain.ErrUserNotFound
	}

	// 2. 验证旧密码是否正确
	if !checkPassword(user, oldPassword) {
		return userDomain.ErrInvalidPassword
	}

	// 3. 校验新密码长度
	if len(newPassword) < 6 {
		return userDomain.ErrPasswordTooShort
	}

	// 4. 加密新密码
	hashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	// 5. 更新用户密码
	user.Password = hashedPassword
	user.UpdatedAt = time.Now()

	// 6. 调用 DAO 层更新
	return u.userRepo.Update(user)
}

// ==================== 私有辅助函数 ====================

// hashPassword 密码加密
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// checkPassword 验证密码
func checkPassword(user *entity.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
