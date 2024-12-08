package Model

type Role struct {
	ID        int    `gorm:"primaryKey;column:id;autoIncrement"`
	RoleTitle string `gorm:"column:roleTitle"`
	QStatus   int    `gorm:"column:qStatus"`
}
