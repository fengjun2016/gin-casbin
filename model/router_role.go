package model

type RouterRole struct {
	RoleId   string `json:"role_id"`
	RouterId string `json:"router_id"`
}

func (this *RouterRole) TableName() string {
	return "router_role"
}
