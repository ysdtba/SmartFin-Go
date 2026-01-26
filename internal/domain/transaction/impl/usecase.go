package impl

import (
	"github.com/shopspring/decimal"

	txRepo "github.com/florentyang/smartfin-go/internal/dao/transaction"
	txDomain "github.com/florentyang/smartfin-go/internal/domain/transaction"
	"github.com/florentyang/smartfin-go/internal/entity"
)

// ==================== UseCase 结构体 ====================

type usecase struct {
	txRepo txRepo.Repo // 依赖 DAO 层接口
}

// ==================== 构造函数 ====================

// NewTransactionDomain 创建 Domain 实例
func NewTransactionDomain(repo txRepo.Repo) txDomain.Domain {
	return &usecase{
		txRepo: repo,
	}
}

// ==================== 业务方法实现 ====================

// Create 创建交易记录
// 核心业务逻辑都在这里：参数校验、金额计算
func (u *usecase) Create(input *txDomain.CreateInput) (*entity.Transaction, error) {

	// ========== 业务规则校验 ==========

	// 1. 校验交易类型：只能是 BUY 或 SELL
	if input.Type != entity.TransactionTypeBuy && input.Type != entity.TransactionTypeSell {
		return nil, txDomain.ErrInvalidType
	}

	// 2. 校验数量：必须大于 0
	if input.Quantity.LessThanOrEqual(decimal.Zero) {
		return nil, txDomain.ErrInvalidQuantity
	}

	// 3. 校验单价：必须大于 0
	if input.Price.LessThanOrEqual(decimal.Zero) {
		return nil, txDomain.ErrInvalidPrice
	}

	// ========== 核心计算 ==========

	// 4. 计算成交总金额：数量 × 单价
	//    使用 decimal 库保证精度，避免浮点数误差
	amount := input.Quantity.Mul(input.Price)

	// ========== 创建实体 ==========

	// 5. 组装 Transaction 实体
	tx := &entity.Transaction{
		UserID:    input.UserID,
		Symbol:    input.Symbol,
		Name:      input.Name,
		Type:      input.Type,
		Quantity:  input.Quantity,
		Price:     input.Price,
		Amount:    amount, // 后端计算的总金额
		Fee:       input.Fee,
		TradeTime: input.TradeTime,
		Notes:     input.Notes,
	}

	// ========== 持久化 ==========

	// 6. 调用 DAO 层存入数据库
	if err := u.txRepo.Create(tx); err != nil {
		return nil, err
	}

	return tx, nil
}
