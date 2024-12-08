package Model

type ProductImage struct {
	ID        int    `gorm:"primaryKey;column:id"`
	ProductID int    `gorm:"primaryKey;column:productId"`
	ImageURL  string `gorm:"column:imageUrl"`
	AddTime   int    `gorm:"column:addTime"`
}
