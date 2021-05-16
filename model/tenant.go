package model

import "fmt"

type Tenant struct {
	Id   string `json:"id" gorm:"type:varchar(20);primaryKey"`     //租户隔离域id
	Name string `json:"name" gorm:"type:varchar(30);unique_index"` //租户隔离域名称
}

func (this *Tenant) String() string {
	return fmt.Sprintf("%d:%s", this.Id, this.Name)
}
