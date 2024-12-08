package Model

type Product struct {
	ID          int    `gorm:"primaryKey;column:id;autoIncrement"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Price       int    `gorm:"column:price"`
	Stock       string `gorm:"column:stock"`
	CategoryID  int    `gorm:"column:categoryId"`
	AddTime     int    `gorm:"column:addTime"`
	UpdateTime  int    `gorm:"column:updateTime"`
	QStatus     int    `gorm:"column:qStatus"`
}
