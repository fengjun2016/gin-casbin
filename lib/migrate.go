package lib

import (
	"gin-casbin/model"
)

func CreateTable() {
	Gorm.AutoMigrate(
		//!!do not delete the line, gen generate code at here
		&model.User{},
		&model.Role{},
		&model.UserRole{},
		&model.Router{},
		&model.RouterRole{},
		&model.Tenant{},
	)
}
