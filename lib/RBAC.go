package lib

import "casbin-with-mysql-go/models"

type RoleRel struct {
	PRole 	string
	Role 	string
}

func (this *RoleRel) String() string {
	return this.PRole+":"+this.Role
}

// 获取角色
func GetRoles(pid int,m *[]*RoleRel,pName string)  {
	proles:=make([]*models.Roles,0)
	Gorm.Where("role_pid=?",pid).Find(&proles)
	if len(proles)==0 {
		return
	}
	for _,item := range proles{
		if pName!="" {
			*m=append(*m,&RoleRel{pName,item.RoleName})
		}
		GetRoles(item.RoleId,m,item.RoleName)
	}
}
