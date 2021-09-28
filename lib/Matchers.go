package lib

import (
	"errors"
)


func init()  {
	E.AddFunction("MethodMatch", func(arguments ...interface{}) (interface{}, error) {
		if len(arguments)==2 {
			key1,key2 := arguments[0].(string),arguments[0].(string)

			return MethodMatch(key1,key2),nil
		}
		return nil,errors.New("MethodMatch error")
	})
}

func MethodMatch(key1, key2 string) bool {
	if key1==key2 {
		return true
	}
	return false
}
