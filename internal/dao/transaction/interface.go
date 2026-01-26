package transaction

import (
	"errors"
	"time"

	"github.com/florentyang/smartfin-go/internal/entity"
)

// ==================== 错误定义 ====================

var (
	ErrTransactionNotFound = errors.New("交易记录不存在")
)

// ==================== 查询条件结构体 ====================

// ListFilter 查询交易列表的筛选条件
type ListFilter struct {
	UserID    uint       // 用户ID（必须）
	Symbol    string     // 股票代码（可选）
	Type      string     // 交易类型（可选）
	StartTime *time.Time // 开始时间（可选）
	EndTime   *time.Time // 结束时间（可选）
	Page      int        // 页码
	PageSize  int        // 每页条数
}

// ==================== 接口定义 ====================
// Domain 层会依赖这个接口

type Repo interface {
	// Create 创建交易记录
	Create(tx *entity.Transaction) error

	// FindByUserID 根据用户ID和筛选条件查询交易列表
	// 返回：交易列表、总条数、错误
	FindByUserID(filter *ListFilter) ([]*entity.Transaction, int64, error)
}
