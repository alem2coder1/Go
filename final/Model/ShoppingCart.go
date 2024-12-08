package Model

type ShoppingCart struct {
	ID      int `gorm:"primaryKey;column:id"`
	UserID  int `gorm:"primaryKey;column:userId"`
	AddTime int `gorm:"column:addTime"`
}
