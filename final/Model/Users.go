package Model

type Users struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"column:name"`
	Age      int    `gorm:"column:age"`
	Surname  string `gorm:"column:surname"`
	Job      string `gorm:"column:job"`
	Birthday string `gorm:"column:birthday"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Status   int    `gorm:"column:status"`
	Role     string `gorm:"column:role"`
}
