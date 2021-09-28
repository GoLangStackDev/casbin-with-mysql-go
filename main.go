package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {
	sub := "lili"
	obj := "/posts"
	act := "GET"

	e,err := casbin.NewEnforcer("resources/model.conf","resources/policy.csv")
	checkError(err)
	ok,err := e.Enforce(sub,obj,act)
	checkError(err)
	if ok {
		fmt.Println("通过！")
	}else{
		fmt.Println("不通过！")
	}
}
// 统一错误检查
func checkError(err error)  {
	if err!=nil {
		println(err.Error())
	}
}
