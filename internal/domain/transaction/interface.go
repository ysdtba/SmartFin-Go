package transaction

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"

	"github.com/florentyang/smartfin-go/internal/entity"
)

// ==================== 错误定义 ====================
// 领域层的业务错误（中文方便调试）

var (
	ErrTransactionNotFound = errors.New("交易记录不存在")
	ErrInvalidType         = errors.New("交易类型无效，必须是 BUY 或 SELL")
	ErrInvalidQuantity     = errors.New("交易数量必须大于 0")
	ErrInvalidPrice        = errors.New("交易单价必须大于 0")
)

// ==================== Domain 输入结构体 ====================
// Service 层通过这些结构体向 Domain 层传递参数

// CreateInput 创建交易的输入参数
type CreateInput struct {
	UserID    uint
	Symbol    string
	Name      string
	Type      string
	Quantity  decimal.Decimal
	Price     decimal.Decimal
	Fee       decimal.Decimal
	TradeTime time.Time
	Notes     string
}

// ==================== Domain 接口定义 ====================
// Service 层会依赖这个接口

type Domain interface {
	// Create 创建交易记录
	// 核心业务逻辑：校验参数、计算总金额、存入数据库
	Create(input *CreateInput) (*entity.Transaction, error)
}
