package lib

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)
var E *casbin.Enforcer

func init()  {
	InitDB()

	adapter, _ := gormadapter.NewAdapterByDB(Gorm)
	e,_ := casbin.NewEnforcer("resources/model.conf", adapter)
	// 必须执行
	e.LoadPolicy()

	E=e
	initPolicy()
}

func initPolicy()  {

	// 加载权限组
	m := make([]*RoleRel,0)
	GetRoles(0,&m,"")
	for _,r := range m{
		E.AddRoleForUser(r.PRole,r.Role)
	}

	// 加载用户和权限
	users := GetUserRoles()
	for _, user := range users {
		E.AddRoleForUser(user.UserName,user.RoleName)
	}
}