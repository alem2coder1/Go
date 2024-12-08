package Model

type Payment struct {
	ID            int    `gorm:"primaryKey;column:id"`
	OrderID       int    `gorm:"primaryKey;column:orderId"`
	Amount        int    `gorm:"column:amount"`
	AddTime       int    `gorm:"column:addTime"`
	PaymentMethod string `gorm:"column:paymentMethod"`
}
