package helpers

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

func UniqueEmail(email string) bool {

	var user models.User

	facades.Orm().Query().Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return true
	}

	return false
}
