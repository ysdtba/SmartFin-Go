package impl

import (
	"gorm.io/gorm"

	txRepo "github.com/florentyang/smartfin-go/internal/dao/transaction"
	"github.com/florentyang/smartfin-go/internal/entity"
)

// ==================== Repository 结构体 ====================

type repository struct {
	db *gorm.DB
}

// ==================== 构造函数 ====================

// NewTransactionRepo 创建 DAO 实例
func NewTransactionRepo(db *gorm.DB) txRepo.Repo {
	return &repository{db: db}
}

// ==================== 接口实现 ====================

// Create 创建交易记录
func (r *repository) Create(tx *entity.Transaction) error {
	return r.db.Create(tx).Error
}
