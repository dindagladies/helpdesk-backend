package models

import (
	"goravel/database/factories"
	"time"

	"github.com/goravel/framework/contracts/database/factory"
)

type User struct {
	// orm.Model
	ID        int       `form:"id" json:"id"`
	Full_Name string    `form:"full_name" json:"full_name"`
	Email     string    `form:"email" json:"email"`
	Password  string    `form:"password" json:"password"`
	Role      string    `form:"role" json:"role"`
	Is_Active bool      `form:"is_active" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *User) TableName() string {
	return "users"
}

func (r *User) Connection() string {
	return "postgresql"
}

func (u *User) Factory() factory.Factory {
	return &factories.UserFactory{}
}
