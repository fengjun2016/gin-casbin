package model

type UserRole struct {
	RoleId string `json:"role_id"`
	UserId string `json:"user_id"`
}

func (this *UserRole) TableName() string {
	return "user_role"
}
