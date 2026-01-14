package impl

import (
	"errors"

	"gorm.io/gorm"

	userRepo "github.com/florentyang/smartfin-go/internal/dao/user"
	"github.com/florentyang/smartfin-go/internal/entity"
)

// ==================== Repository 结构体（依赖聚合） ====================

type repository struct {
	db *gorm.DB
	// 以后可以加更多依赖：
	// cache *redis.Client
}

// ==================== 构造函数 ====================

// NewUserRepo 创建 DAO 实例
// 返回接口类型，隐藏实现细节
func NewUserRepo(db *gorm.DB) userRepo.Repo {
	return &repository{db: db}
}

// ==================== 接口实现 ====================

// Create 创建用户
func (r *repository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

// GetByID 按 ID 查找用户
func (r *repository) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userRepo.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// GetByUsername 按用户名查找用户
func (r *repository) GetByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userRepo.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail 按邮箱查找用户
func (r *repository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userRepo.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (r *repository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

// ExistsByUsername 检查用户名是否已存在
func (r *repository) ExistsByUsername(username string) bool {
	var count int64
	r.db.Model(&entity.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

// ExistsByEmail 检查邮箱是否已存在
func (r *repository) ExistsByEmail(email string) bool {
	var count int64
	r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

