package factories

type UserFactory struct {
}

// Definition Define the model's default state.
func (f *UserFactory) Definition() map[string]any {
	return map[string]any{
		"full_name": "John Doe",
		"email":     "john@mail.com",
		"password":  "123456",
		"role":      "admin",
		"is_active": true,
	}
}
