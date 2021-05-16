package lib

import (
	"github.com/casbin/casbin/v2"

	"github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Header.Get("token") == "" {
			context.AbortWithStatusJSON(400, gin.H{"message": "token required"})
		} else {
			context.Set("user_name", context.Request.Header.Get("token"))
			context.Next()
		}
	}
}

func RBAC() gin.HandlerFunc {
	adapter, err := gormadapter.NewAdapterByDB(Gorm)
	if err != nil {
		log.Fatal(err)
	}

	e, err := casbin.NewEnforcer("resources/model.conf", adapter)
	if err != nil {
		log.Fatal(err)
	}

	err = e.LoadPolicy()
	if err != nil {
		log.Fatal(err)
	}

	return func(context *gin.Context) {
		user, _ := context.Get("user_name")
		domain := context.Param("domain")
		//租户 后面的 才是 casbin 中的 uri
		uri := strings.TrimPrefix(context.Request.RequestURI, "/"+domain) // 获取租户后面的 路径 参数 /depts
		access, err := e.Enforce(user, domain, uri, context.Request.Method)
		if err != nil || !access {
			context.AbortWithStatusJSON(403, gin.H{"message": "forbidden"})
		} else {
			context.Next()
		}
	}
}

func Middlewares() (fs []gin.HandlerFunc) {
	fs = append(fs, CheckLogin(), RBAC())
	return
}
