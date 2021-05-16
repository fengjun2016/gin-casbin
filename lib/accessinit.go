package lib

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"log"
)

var E *casbin.Enforcer

func init() {
	initDB()
	adapter, err := gormadapter.NewAdapterByDB(Gorm)
	if err != nil {
		log.Fatal(err)
	}

	CreateTable()
	e, err := casbin.NewEnforcer("resources/model_t.conf", adapter)
	if err != nil {
		log.Fatal(err)
	}

	err = e.LoadPolicy()
	if err != nil {
		log.Fatal(err)
	}

	E = e
	// initPolicy()           //不带租户
	initPolicyWithDomain() //带租户隔离域
}

//从我们的库里初始化 策略数据 现在是不带租户的 权限初始化 策略数据
func initPolicy() {
	// E.AddPolicy("member", "/depts", "GET")  // p数据
	// E.AddPolicy("admin", "/depts", "POST")  // p数据
	// E.AddRoleForUser("zhangshan", "member") // g数据

	//获取所有角色 以及对应
	m := make([]*RoleRel, 0)
	GetRoles("", &m, "")

	//遍历 然后初始化策略数据
	//初始化角色权限
	for _, r := range m {
		_, err := E.AddRoleForUser(r.PRole, r.Role)
		if err != nil {
			log.Fatal(err)
		}
	}

	//获取用户角色权限
	//初始化用户角色权限
	userRoles := GetUserRoles()
	for _, ur := range userRoles {
		_, err := E.AddRoleForUser(ur.UserName, ur.RoleName) //前项继承后项
		if err != nil {
			log.Fatal(err)
		}
	}

	//初始化 路由角色
	routerRoles := GetRouterRoles()
	for _, rr := range routerRoles {
		_, err := E.AddPolicy(rr.Name, rr.Uri, rr.Method)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//租户 初始化
func initPolicyWithDomain() {
	//下面 这部分是初始化 角色 关系
	//拼凑出这种格式
	//g, deptadmin, deptupdater, domain1
	//g, deptupdater, deptselecter, domain2
	//其中 deptselecter 权限最低, 然后是deptupdater, 最后是deptadmin

	roles := GetRolesWithDomain() //获取角色对应
	for _, r := range roles {
		_, err := E.AddRoleForUserInDomain(r.PRole, r.Role, r.Domain) //看这一句, 加了domain参数
		if err != nil {
			log.Fatal(err)
		}
	}

	////// 初始化用户角色， 格式和上面一样
	userRoles := GetUserRolesWithDomain()
	for _, ur := range userRoles {
		// 这里也做了改变, 增加了domain参数
		_, err := E.AddRoleForUserInDomain(ur.UserName, ur.RoleName, ur.TenantName)
		if err != nil {
			log.Fatal(err)
		}
	}

	//初始化 路由角色对应关系
	//格式 p deptselecter  domain1 /depts  GET
	routerRoles := GetRouterRolesWithDomain()
	for _, rr := range routerRoles {
		_, err := E.AddPolicy(rr.RoleName, rr.TenantName, rr.Uri, rr.Method)
		if err != nil {
			log.Fatal(err)
		}
	}
}
