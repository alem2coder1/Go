package user

type Role struct {
	ID        int    `gorm:"primaryKey"`
	RoleTitle string `gorm:"column:roleTitle"`
	Status    byte   `gorm:"column:qStatus"`
}
