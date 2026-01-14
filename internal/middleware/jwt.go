package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/florentyang/smartfin-go/pkg/jwt"
	"github.com/florentyang/smartfin-go/pkg/response"
)

// JWTAuth JWT 鉴权中间件
// 验证请求头中的 Token，并将用户信息存入 Context
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从 Header 获取 Token
		// 格式：Authorization: Bearer eyJhbGci...
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "请先登录")
			c.Abort()
			return
		}

		// 2. 解析 Bearer Token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "Token 格式错误")
			c.Abort()
			return
		}
		tokenString := parts[1]

		// 3. 验证 Token（调用 jwt.go 的工具函数）
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			response.Unauthorized(c, "Token 无效或已过期")
			c.Abort()
			return
		}

		// 4. 将用户信息存入 Context，供后续 Controller 使用
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		// 5. 继续执行后续 Handler
		c.Next()
	}
}
