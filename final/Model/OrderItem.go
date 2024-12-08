package Model

type OrderItem struct {
	ID        int `gorm:"primaryKey;column:id"`
	OrderID   int `gorm:"column:orderId"`
	ProductID int `gorm:"column:productId"`
	Quantity  int `gorm:"column:quantity"`
	Price     int `gorm:"column:price"`
}
