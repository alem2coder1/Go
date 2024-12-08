package Model

type CartItem struct {
	ID        int `gorm:"primaryKey;column:id"`
	CardID    int `gorm:"column:cardId"`
	ProductID int `gorm:"column:productId"`
	Quantity  int `gorm:"column:quantity"`
}
