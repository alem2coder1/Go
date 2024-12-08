package Model

type AuditLog struct {
	ID        int    `gorm:"primaryKey;column:id"`
	UserID    int    `gorm:"primaryKey;column:userId"`
	Action    string `gorm:"column:action"`
	Timestamp int    `gorm:"column:timestamp"`
}
