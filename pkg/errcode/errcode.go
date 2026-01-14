package errcode

// 通用错误码 (1000-1999)
const (
	Success         = 0
	ServerError     = 1001
	InvalidParams   = 1002
	NotFound        = 1003
	TooManyRequests = 1004
)

// 用户模块错误码 (2000-2999)
const (
	UserNotFound      = 2001
	UserAlreadyExists = 2002
	PasswordError     = 2003
	TokenInvalid      = 2004
	TokenExpired      = 2005
)

// 资产模块错误码 (3000-3999)
const (
	AssetNotFound       = 3001
	InsufficientBalance = 3002
	AssetAlreadyExists  = 3003
)

// 交易模块错误码 (4000-4999)
const (
	TransactionNotFound = 4001
	TransactionFailed   = 4002
	InvalidAmount       = 4003
	InvalidQuantity     = 4004
)

// 行情模块错误码 (5000-5999)
const (
	StockNotFound   = 5001
	QuoteFetchError = 5002
	MarketClosed    = 5003
)

// AI 模块错误码 (6000-6999)
const (
	AIServiceError = 6001
	QueryTooLong   = 6002
	RAGIndexError  = 6003
)

// ErrMsg 错误码对应的消息
var ErrMsg = map[int]string{
	Success:         "操作成功",
	ServerError:     "服务器内部错误",
	InvalidParams:   "参数错误",
	NotFound:        "资源不存在",
	TooManyRequests: "请求过于频繁",

	UserNotFound:      "用户不存在",
	UserAlreadyExists: "用户已存在",
	PasswordError:     "密码错误",
	TokenInvalid:      "Token 无效",
	TokenExpired:      "Token 已过期",

	AssetNotFound:       "资产不存在",
	InsufficientBalance: "余额不足",
	AssetAlreadyExists:  "资产已存在",

	TransactionNotFound: "交易记录不存在",
	TransactionFailed:   "交易失败",
	InvalidAmount:       "金额无效",
	InvalidQuantity:     "数量无效",

	StockNotFound:   "股票不存在",
	QuoteFetchError: "行情获取失败",
	MarketClosed:    "市场已休市",

	AIServiceError: "AI 服务异常",
	QueryTooLong:   "查询内容过长",
	RAGIndexError:  "知识库索引错误",
}

// GetMsg 根据错误码获取错误消息
func GetMsg(code int) string {
	if msg, ok := ErrMsg[code]; ok {
		return msg
	}
	return "未知错误"
}

