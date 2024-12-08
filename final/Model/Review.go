package Model

type Review struct {
	ID        int    `gorm:"primaryKey;column:id"`
	ProductID int    `gorm:"primaryKey;column:productId"`
	UserID    int    `gorm:"primaryKey;column:userId"`
	Rating    string `gorm:"column:rating"`
	Comment   string `gorm:"column:comment"`
	AddTime   int    `gorm:"column:addTime"`
}
