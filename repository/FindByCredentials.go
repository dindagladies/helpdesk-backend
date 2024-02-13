package repository

import (
	"errors"

	"heldesk-api/database"
	"heldesk-api/model"
)

func FindByCredentials(email, password string) (*model.User, error) {
	// Here you would query your database for the user with the given email

	/*
	* Validate email
	 */

	db := database.DB.Db
	var user model.User
	db.Find(&user, "email = ?", email)

	if user.Email == "" {
		return nil, errors.New("email not found")
	}

	/*
	* Validate password
	 */
	credentialError := user.CheckPassword(password)
	if credentialError != nil {
		return nil, errors.New("password does not match")
	}

	return &user, nil
}
