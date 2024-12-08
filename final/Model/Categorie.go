package Model

type Categories struct {
	ID          int    `gorm:"primaryKey;column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	QStatus     int    `gorm:"column:qStatus"`
}
