package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/florentyang/smartfin-go/internal/dto"
	"github.com/florentyang/smartfin-go/internal/service"
	"github.com/florentyang/smartfin-go/pkg/response"
)

// ==================== 接口定义 ====================

type TransactionController interface {
	Create(c *gin.Context) // 创建交易
	List(c *gin.Context)   // 查询交易列表
}

// ==================== 结构体 ====================

type transactionController struct {
	txService service.TransactionService
}

// ==================== 构造函数 ====================

func NewTransactionController(txService service.TransactionService) TransactionController {
	return &transactionController{txService: txService}
}

// ==================== 接口实现 ====================

// Create 创建交易记录
// POST /api/v1/transactions
// 请求体：{ symbol, name, type, quantity, price, fee, trade_time, notes }
func (ctrl *transactionController) Create(c *gin.Context) {
	// 1. 从 JWT 中间件获取用户ID（确保用户已登录）
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	// 2. 绑定请求参数（JSON → DTO）
	var req dto.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 3. 调用 Service 层处理业务
	tx, err := ctrl.txService.Create(userID.(uint), &req)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	// 4. 返回创建成功的交易记录
	response.Success(c, tx)
}

// List 查询交易列表
// GET /api/v1/transactions
// Query 参数：page, page_size, symbol, type, start_date, end_date
func (ctrl *transactionController) List(c *gin.Context) {
	// 1. 从 JWT 中间件获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	// 2. 绑定 Query 参数（URL → DTO）
	var req dto.ListTransactionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 3. 调用 Service 层查询
	result, err := ctrl.txService.List(userID.(uint), &req)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	// 4. 返回分页数据
	response.Success(c, result)
}
