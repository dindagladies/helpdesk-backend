package controllers

import (
	"helpdesk/api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) Index(c *gin.Context) {
	var (
		users  []structs.User
		result gin.H
	)

	idb.DB.Find(&users)
	if len(users) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": users,
			"count":  0,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) Show(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)

	id := c.Param("id")
	print(id)
	err := idb.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": user,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) Store(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)

	full_name := c.PostForm("full_name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role := c.PostForm("role")
	// is_active := c.PostForm("is_active")

	user.Full_Name = full_name
	user.Email = email
	user.Password = password
	user.Role = role
	// user.Is_Active = is_active

	idb.DB.Create(&user)
	result = gin.H{
		"result": user,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) Update(c *gin.Context) {
	id := c.Param("id")
	full_name := c.PostForm("full_name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role := c.PostForm("role")
	// is_active := c.PostForm("is_active")
	// updated_at := c.PostForm("updated_at")

	var (
		user    structs.User
		newUser structs.User
		result  gin.H
	)

	// get data by id
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newUser.Full_Name = full_name
	newUser.Email = email
	newUser.Password = password
	newUser.Role = role
	// newUser.Updated_At = updated_at

	// update data
	err = idb.DB.Model(&user).Updates(newUser).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) Delete(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)

	// get data by id
	id := c.Param("id")
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	// delete data
	err = idb.DB.Delete(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
