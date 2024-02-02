package handler

import (
	"heldesk-api/database"
	"heldesk-api/model"

	// "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User

	// find all users
	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Users not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Users found",
		"data":    users,
	})

}

func GetUser(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")
	var user model.User

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User found",
		"data":    user,
	})

}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	// store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create user",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "User has created",
		"data":    user,
	})
}

func updateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Username string `json:"username"`
	}

	db := database.DB.Db
	var user model.User
	id := c.Params("id")

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	user.Username = updateUserData.Username

	db.Save(&user)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "users found",
		"data":    user,
	})
}

// func DeleteUserByID(c *fiber.Ctx) error {
// 	db := database.DB.Db
// 	var user model.User

// 	id := c.Params("id")
// 	db.Find(&user, "id = ?", id)

// 	if user.ID == uuid.Nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status":  "error",
// 			"message": "User not found",
// 			"data":    nil,
// 		})
// 	}
// }
