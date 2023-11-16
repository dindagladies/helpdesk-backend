package structs

import (
	"time"
)

type User struct {
	// gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id" form:"id"`
	Full_Name string    `json:"full_name" form:"full_name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Role      string    `json:"role" form:"role"`
	Is_Active bool      `json:"is_active" form:"is_active"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
}
