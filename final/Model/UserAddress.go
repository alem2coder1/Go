package Model

type UserAddress struct {
	ID      int    `gorm:"primaryKey;column:id"`
	UserID  int    `gorm:"primaryKey;column:userId"`
	Street  string `gorm:"column:street"`
	City    string `gorm:"column:city"`
	State   string `gorm:"column:state"`
	ZipCode string `gorm:"column:zipCode"`
}
