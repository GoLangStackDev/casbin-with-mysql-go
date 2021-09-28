package models

import "fmt"

type Users struct {
	UserId 		int			`gorm:"primaryKey"`
	UserName 	string
	RoleName 	string
}

func (this *Users) String() string {
	return fmt.Sprintf("%s:%s",this.UserName,this.RoleName)
}