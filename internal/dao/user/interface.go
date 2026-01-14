package user

import (
	"errors"

	"github.com/florentyang/smartfin-go/internal/entity"
)

// ==================== 错误定义 ====================
// DAO 层的错误，供上层判断使用

var (
	ErrUserNotFound = errors.New("用户不存在")
)

// ==================== 接口定义 ====================
// Domain 层会依赖这个接口，而不是具体实现

type Repo interface {
	// Create 创建用户
	Create(user *entity.User) error

	// GetByID 按 ID 查找用户
	GetByID(id uint) (*entity.User, error)

	// GetByUsername 按用户名查找用户
	GetByUsername(username string) (*entity.User, error)

	// GetByEmail 按邮箱查找用户
	GetByEmail(email string) (*entity.User, error)

	// Update 更新用户信息
	Update(user *entity.User) error

	// ExistsByUsername 检查用户名是否已存在
	ExistsByUsername(username string) bool

	// ExistsByEmail 检查邮箱是否已存在
	ExistsByEmail(email string) bool
}

