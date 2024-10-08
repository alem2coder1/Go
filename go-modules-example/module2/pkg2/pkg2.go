package pkg2

import "fmt"

func HelloFromPkg2() {
	fmt.Println("from from module 2!")
}

type Users struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"column:name"`
	Surname  string `gorm:"column:surname"`
	Age      int    `gorm:"column:age"`
	Birthday string `gorm:"column:birthday"`
	Job      string `gorm:"column:job"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}
