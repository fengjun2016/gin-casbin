package lib

import (
	"gin-casbin/model"
	"log"
)

type RoleRel struct {
	PRole  string `json:"p_role"` //父级角色
	Role   string `json:"role"`   //子级角色
	Domain string `json:"domain"` //租户的域
}

func (this *RoleRel) String() string {
	return this.PRole + ":" + this.Role + ":" + this.Domain
}

func AllTenants() (ret []*model.Tenant) { //获取所有租户
	db := Gorm.Table("tenants").Find(&ret)
	if db.Error != nil {
		log.Fatal(db.Error)
	}

	return
}

//获取角色 ---不带租户
func GetRoles(pid string, m *[]*RoleRel, pname string) {
	proles := make([]*model.Role, 0)
	Gorm.Table("role").
		Where("p_id = ?", pid).
		Find(&proles)

	if len(proles) == 0 {
		return
	}

	for _, item := range proles {
		if pname != "" {
			*m = append(*m, &RoleRel{pname, item.Name, ""})
		}
		GetRoles(item.Id, m, item.Name)
	}
}

//获取角色 -- 带租户
func GetRolesWithDomain() []*RoleRel {
	ts := AllTenants() //获取 所有 租户
	roleRels := make([]*RoleRel, 0)
	for _, t := range ts { //遍历租户
		t_roleRels := make([]*RoleRel, 0)
		getRolesWithDomain("", &t_roleRels, "", t)
		roleRels = append(roleRels, t_roleRels...)
	}

	return roleRels
}

func getRolesWithDomain(pid string, m *[]*RoleRel, pname string, t *model.Tenant) {
	proles := make([]*model.Role, 0)
	//注意这里, 根据每个租户ID进行获取
	Gorm.Where("p_id = ? and tenant_id = ?", pid, t.Id).Find(&proles)
	if len(proles) == 0 {
		return
	}

	for _, item := range proles {
		if pname != "" {
			*m = append(*m, &RoleRel{pname, item.Name, t.Name})
		}
		getRolesWithDomain(item.Id, m, item.Name, t)
	}
}

//获取用户和角色对应关系
func GetUserRoles() (users []*model.User) {
	Gorm.Table("user a, user_role b, role c").
		Select("a.user_name,c.role_name").
		Where("a.id = b.user_id and b.role_id = c.id").
		Find(&users)

	return
}

//获取用户和角色对应关系 ---带租户
func GetUserRolesWithDomain() (users []*model.User) {
	Gorm.Table("user a, user_role b, role c, tenants d").
		Select("a.user_name as user_name, c.name as role_name, d.name as tenant_name").
		Where("a.id = b.user_id and b.role_id = c.id and c.tenant_id = d.id").
		Find(&users)

	return
}

//获取路由和角色对应关系
func GetRouterRoles() (routers []*model.Router) {
	Gorm.Table("router a, router_role b, role c").
		Select("a.uri, a.method, c.name").
		Where("a.id = b.user_id and b.role_id = c.id").
		Find(&routers)

	return
}

//获取路由和角色对应关系  ---带租户
func GetRouterRolesWithDomain() (routers []*model.Router) {
	Gorm.Table("router a, router_role b, role c, tenants d").
		Select("a.uri, a.method, c.name, d.name").
		Where("a.id = b.router_id and b.role_id = c.id and c.tenant_id = d.id").
		Find(&routers)

	return
}
