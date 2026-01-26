package transaction

import (
	"errors"

	"github.com/florentyang/smartfin-go/internal/entity"
)

// ==================== 错误定义 ====================

var (
	ErrTransactionNotFound = errors.New("交易记录不存在")
)

// ==================== 接口定义 ====================
// Domain 层会依赖这个接口

type Repo interface {
	// Create 创建交易记录
	Create(tx *entity.Transaction) error
}
