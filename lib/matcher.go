package lib

import (
	"fmt"
	"strings"
)

func init() {
	E.AddFunction("methodMatch", func(argus ...interface{}) (i interface{}, e error) {
		if len(argus) == 2 {
			k1, k2 := argus[0].(string), argus[1].(string)
			return bool(MethodMatch(k1, k2)), nil
		}

		return nil, fmt.Errorf("method not matched")
	})
	E.AddFunction("isSuper", func(arguments ...interface{}) (i interface{}, e error) {
		if len(arguments) == 1 {
			user := arguments[0].(string)

			return IsSuperAdmin(user), nil
		}

		return nil, fmt.Errorf("methodMatch error")
	})
}

func MethodMatch(key1 string, key2 string) bool {
	ks := strings.Split(key2, " ")
	for _, s := range ks {
		if s == key1 {
			return true
		}
	}

	return false
}

//后面可以改成从数据库取
var ADMINS = []string{"admin", "root"}

func IsSuperAdmin(userName string) bool {
	for _, user := range ADMINS {
		if user == userName {
			return true
		}
	}

	return false
}
