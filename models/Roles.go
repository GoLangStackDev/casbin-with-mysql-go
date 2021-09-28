package models

import "fmt"

type Roles struct {
	RoleId 		int 	`gorm:"primaryKey"`
	RoleName 	string
	RolePid 	int
	RoleComment string
}

func (this *Roles) String() string {
	return fmt.Sprintf("ID:%d 角色名:%s",this.RoleId,this.RoleName)
}