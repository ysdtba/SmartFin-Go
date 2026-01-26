package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

// ================== 请求 DTO ==================

// CreateTransactionRequest 创建交易请求
// 前端传来的 JSON 会自动映射到这个结构体
type CreateTransactionRequest struct {
	Symbol    string          `json:"symbol" binding:"required"`              // 股票代码，如 AAPL
	Name      string          `json:"name"`                                   // 股票名称（可选）
	Type      string          `json:"type" binding:"required,oneof=BUY SELL"` // 交易类型：必须是 BUY 或 SELL
	Quantity  decimal.Decimal `json:"quantity" binding:"required"`            // 交易数量
	Price     decimal.Decimal `json:"price" binding:"required"`               // 成交单价
	Fee       decimal.Decimal `json:"fee"`                                    // 手续费（可选，默认0）
	TradeTime string          `json:"trade_time" binding:"required"`          // 交易时间，ISO 8601 格式：2024-01-15T10:30:00Z
	Notes     string          `json:"notes"`                                  // 备注（可选）
}

// ListTransactionRequest 查询交易列表请求
// 使用 form 标签绑定 Query 参数
type ListTransactionRequest struct {
	// ===== 分页参数 =====
	Page     int `form:"page"`      // 页码，默认 1
	PageSize int `form:"page_size"` // 每页条数，默认 20

	// ===== 筛选参数 =====
	Symbol    string `form:"symbol"`     // 按股票代码筛选（可选）
	Type      string `form:"type"`       // 按交易类型筛选：BUY/SELL（可选）
	StartDate string `form:"start_date"` // 开始日期：2024-01-01（可选）
	EndDate   string `form:"end_date"`   // 结束日期：2024-12-31（可选）
}

// ================== 响应 DTO ==================

// TransactionResponse 交易响应
// 返回给前端的数据结构
type TransactionResponse struct {
	ID        uint            `json:"id"`
	Symbol    string          `json:"symbol"`
	Name      string          `json:"name"`
	Type      string          `json:"type"`
	Quantity  decimal.Decimal `json:"quantity"`
	Price     decimal.Decimal `json:"price"`
	Amount    decimal.Decimal `json:"amount"` // 总金额（后端计算）
	Fee       decimal.Decimal `json:"fee"`
	TradeTime time.Time       `json:"trade_time"`
	Notes     string          `json:"notes"`
	CreatedAt time.Time       `json:"created_at"`
}

// ListTransactionResponse 分页列表响应
type ListTransactionResponse struct {
	Total    int64                  `json:"total"`     // 总条数
	Page     int                    `json:"page"`      // 当前页码
	PageSize int                    `json:"page_size"` // 每页条数
	List     []*TransactionResponse `json:"list"`      // 数据列表
}
