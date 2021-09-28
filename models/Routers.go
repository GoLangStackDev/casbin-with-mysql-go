package models

import "fmt"

type Routers struct {
	RId 		int 		`gorm:"column:r_id;primaryKey"`
	RName 		string 		`gorm:"column:r_name"`
	RUri 		string
	RMethod 	string
	RStatus 	int
	RoleName 	string
}
func (this *Routers) String() string {
	return fmt.Sprintf("ID:%s 角色名:%s",this.RName,this.RoleName)
}