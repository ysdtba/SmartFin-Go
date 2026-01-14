package entity

import "time"

// User 用户实体（对应数据库表）
// 放在 entity 包，避免循环依赖
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"not null;uniqueIndex;size:50"`
	Email     string    `gorm:"not null;uniqueIndex;size:100"`
	Password  string    `gorm:"not null;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
