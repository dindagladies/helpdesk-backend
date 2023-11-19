package rules

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type Unique struct {
}

// Signature The name of the rule.
func (receiver *Unique) Signature() string {
	return "unique"
}

// Passes Determine if the validation rule passes.
func (receiver *Unique) Passes(data validation.Data, val any, options ...any) bool {
	// unique email
	var user models.User
	facades.Orm().Query().Where("email = ?", val).First(&user)
	if user.ID == 0 {
		return false
	}

	return true
}

// Message Get the validation error message.
func (receiver *Unique) Message() string {
	return ""
}
