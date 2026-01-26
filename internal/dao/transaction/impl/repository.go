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

// FindByUserID 根据用户ID和筛选条件查询交易列表
// 支持：分页、按股票代码筛选、按交易类型筛选、按日期范围筛选
func (r *repository) FindByUserID(filter *txRepo.ListFilter) ([]*entity.Transaction, int64, error) {
	var txList []*entity.Transaction
	var total int64

	// ===== 构建基础查询（必须按用户ID筛选） =====
	query := r.db.Model(&entity.Transaction{}).Where("user_id = ?", filter.UserID)

	// ===== 动态添加筛选条件 =====

	// 按股票代码筛选
	if filter.Symbol != "" {
		query = query.Where("symbol = ?", filter.Symbol)
	}

	// 按交易类型筛选
	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}

	// 按日期范围筛选（开始时间）
	if filter.StartTime != nil {
		query = query.Where("trade_time >= ?", filter.StartTime)
	}

	// 按日期范围筛选（结束时间）
	if filter.EndTime != nil {
		query = query.Where("trade_time < ?", filter.EndTime)
	}

	// ===== 先查询总数（分页前） =====
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ===== 分页 + 排序 + 查询数据 =====
	offset := (filter.Page - 1) * filter.PageSize

	err := query.
		Order("trade_time DESC"). // 默认按交易时间倒序
		Limit(filter.PageSize).
		Offset(offset).
		Find(&txList).Error

	if err != nil {
		return nil, 0, err
	}

	return txList, total, nil
}
