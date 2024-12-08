package Model

type Order struct {
	ID          int `gorm:"primaryKey;column:id;autoIncrement"`
	UserID      int `gorm:"column:user_id"`
	AddTime     int `gorm:"column:addTime"`
	QStatus     int `gorm:"column:qStatus"`
	TotalAmount int `gorm:"column:totalAmount"`
}
