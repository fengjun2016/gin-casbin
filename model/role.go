package model

import "fmt"

type Role struct {
	Id         string `json:"id" gorm:"type:varchar(30);primaryKey"`
	Name       string `json:"name" gorm:"type:varchar(40)"`            //角色名称
	PId        string `json:"p_id" gorm:"type:varchar(40)"`            //父角色id
	Comment    string `json:"comment"  gorm:"type:text"`               //角色备注
	TenantId   string `json:"tenant_id" gorm:"type:varchar(20);index"` //租户隔离域id
	TenantName string `json:"tenant_name"`                             //租户隔离域名称
}

func (this *Role) TableName() string {
	return "role"
}

func (this *Role) String() string {
	return fmt.Sprintf("ID:%d 角色名:%s", this.Id, this.Name)
}
