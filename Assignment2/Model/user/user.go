package user

type Users struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"column:name"`
	Role     string `gorm:"column:role"`
	Surname  string `gorm:"column:surname"`
	Age      int    `gorm:"column:age"`
	Birthday string `gorm:"column:birthday"`
	Job      string `gorm:"column:job"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Status   byte   `gorm:"column:status"`
}
