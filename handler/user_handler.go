package handler

import (
	"heldesk-api/database"
	"heldesk-api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User

	// find all users
	db.Find(&users)

	if len(users) == 0 {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Users is empty.",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Users found successfully.",
		"data":    users,
		"meta": fiber.Map{
			"current_page": 0,
			"last_page":    0,
			"per_page":     0,
			"total":        len(users),
		},
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

	// TODO: remove password from response (by hide user gorm)
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User found successfully.",
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
		"message": "User created successfully.",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Full_Name string `json:"full_name"`
		Role      string `json:"role"`
		Is_Active bool   `json:"is_active"`
		Password  string `json:"password"`
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

	user.Full_Name = updateUserData.Full_Name
	user.Role = updateUserData.Role
	user.Is_Active = updateUserData.Is_Active
	user.HashPassword(updateUserData.Password)

	db.Save(&user)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User updated successfully",
		"data":    user,
	})
}

func DeleteUserByID(c *fiber.Ctx) error {
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

	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete user",
			"data":    nil,
		})
	}

	return c.Status(202).JSON(fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
		"data":    nil,
	})

}
