package Model

type Session struct {
	ID        int    `gorm:"primaryKey;column:id"`
	UserID    int    `gorm:"primaryKey;column:userId"`
	AddTime   int    `gorm:"column:addTime"`
	ExpiresAt string `gorm:"column:expires_at"`
}
