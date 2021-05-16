package model

import "fmt"

type User struct {
	Id         string `json:"id" gorm:"primaryKey"`               //用户id
	UserName   string `json:"user_name" gorm:"type:varchar(20);"` //用户名称
	RoleName   string `json:"role_name" gorm:"type:varchar(20);"` //角色名称
	TenantId   string `json:"tenant_id" gorm:"type:varchar(20);"` //租户id
	TenantName string `json:"tenant_name"`                        //租户域名名称
}

func (this *User) TableName() string {
	return "user"
}

func (this *User) String() string {
	return fmt.Sprintf("%s-%s", this.UserName, this.RoleName)
}
