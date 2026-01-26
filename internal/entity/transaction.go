package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// Transaction 交易记录实体（对应数据库表 transactions）
// 记录用户的每一笔买入/卖出操作
type Transaction struct {
	ID        uint            `gorm:"primaryKey"`                   // 主键ID
	UserID    uint            `gorm:"not null;index"`               // 用户ID（关联 users 表）
	Symbol    string          `gorm:"not null;size:20;index"`       // 股票代码，如 AAPL、TSLA
	Name      string          `gorm:"size:100"`                     // 股票名称，如 Apple Inc.
	Type      string          `gorm:"not null;size:10"`             // 交易类型：BUY（买入）/ SELL（卖出）
	Quantity  decimal.Decimal `gorm:"type:decimal(18,4);not null"`  // 交易数量（用 decimal 防止精度丢失）
	Price     decimal.Decimal `gorm:"type:decimal(18,4);not null"`  // 成交单价
	Amount    decimal.Decimal `gorm:"type:decimal(18,4);not null"`  // 成交总金额 = Quantity × Price
	Fee       decimal.Decimal `gorm:"type:decimal(18,4);default:0"` // 手续费
	TradeTime time.Time       `gorm:"not null;index"`               // 交易时间（用户输入的实际成交时间）
	Notes     string          `gorm:"size:500"`                     // 备注
	CreatedAt time.Time       `gorm:"autoCreateTime"`               // 记录创建时间（系统自动）
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`               // 记录更新时间（系统自动）
}

// 交易类型常量
const (
	TransactionTypeBuy  = "BUY"  // 买入
	TransactionTypeSell = "SELL" // 卖出
)
