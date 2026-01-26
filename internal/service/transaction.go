package service

import (
	"time"

	txDomain "github.com/florentyang/smartfin-go/internal/domain/transaction"
	"github.com/florentyang/smartfin-go/internal/dto"
	"github.com/florentyang/smartfin-go/internal/entity"
)

// ==================== 接口定义 ====================
// Controller 层会使用这个接口

type TransactionService interface {
	Create(userID uint, req *dto.CreateTransactionRequest) (*dto.TransactionResponse, error)
}

// ==================== 接口实现 ====================

type transactionService struct {
	txDomain txDomain.Domain // 依赖 Domain 层接口
}

// NewTransactionService 创建 Service 实例
func NewTransactionService(txDomain txDomain.Domain) TransactionService {
	return &transactionService{
		txDomain: txDomain,
	}
}

// Create 创建交易记录
// Service 层职责：
// 1. 解析时间字符串
// 2. 调用 Domain 层
// 3. Entity → DTO 转换
func (s *transactionService) Create(userID uint, req *dto.CreateTransactionRequest) (*dto.TransactionResponse, error) {
	// 1. 解析交易时间（字符串 → time.Time）
	//    前端传 ISO 8601 格式："2024-01-15T10:30:00Z"
	tradeTime, err := time.Parse(time.RFC3339, req.TradeTime)
	if err != nil {
		return nil, err
	}

	// 2. 调用 Domain 层处理核心业务
	tx, err := s.txDomain.Create(&txDomain.CreateInput{
		UserID:    userID,
		Symbol:    req.Symbol,
		Name:      req.Name,
		Type:      req.Type,
		Quantity:  req.Quantity,
		Price:     req.Price,
		Fee:       req.Fee,
		TradeTime: tradeTime,
		Notes:     req.Notes,
	})
	if err != nil {
		return nil, err
	}

	// 3. Entity → DTO 转换
	return txEntityToDTO(tx), nil
}

// ==================== 私有辅助函数 ====================

// txEntityToDTO 将 Transaction Entity 转换为 DTO
func txEntityToDTO(tx *entity.Transaction) *dto.TransactionResponse {
	return &dto.TransactionResponse{
		ID:        tx.ID,
		Symbol:    tx.Symbol,
		Name:      tx.Name,
		Type:      tx.Type,
		Quantity:  tx.Quantity,
		Price:     tx.Price,
		Amount:    tx.Amount,
		Fee:       tx.Fee,
		TradeTime: tx.TradeTime,
		Notes:     tx.Notes,
		CreatedAt: tx.CreatedAt,
	}
}
