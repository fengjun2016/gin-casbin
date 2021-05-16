package model

type Router struct {
	ID         string `json:"id" gorm:"type:varchar(30);primaryKey"` //id 主键
	Name       string `json:"name" gorm:"type:varchar(30);"`         //路由名称
	Uri        string `json:"uri" gorm:"type:text"`                  //路由路径
	Method     string `json:"method"`                                //路由请求方法
	RoleName   string `json:"role_name"`                             //拥有者的角色名称
	TenantName string `json:"tenant_names"`                          //租户域名
}

func (this *Router) TableName() string {
	return "router"
}
