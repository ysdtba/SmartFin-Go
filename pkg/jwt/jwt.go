package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 配置常量（生产环境应该放到配置文件）
var (
	SecretKey   = []byte("smartfin-secret-key-2026") // JWT 签名密钥
	TokenExpiry = 30 * time.Minute                   // Token 有效期：30分钟
)

// 错误定义
var (
	ErrInvalidToken = errors.New("无效的 Token")
	ErrExpiredToken = errors.New("Token 已过期")
)

// Claims 自定义 JWT 载荷
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT Token
// 参数：用户ID、用户名
// 返回：token 字符串、过期时间戳、错误
func GenerateToken(userID uint, username string) (string, int64, error) {
	// 计算过期时间
	expiresAt := time.Now().Add(TokenExpiry)

	// 创建 Claims
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),  // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			Issuer:    "smartfin-go",                  // 签发者
		},
	}

	// 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名生成字符串
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiresAt.Unix(), nil
}

// ParseToken 解析 JWT Token
// 参数：token 字符串
// 返回：Claims（包含 UserID）、错误
func ParseToken(tokenString string) (*Claims, error) {
	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	// 验证并提取 Claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
