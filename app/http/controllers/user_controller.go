package controllers

import (
	"goravel/app/helpers"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/validation"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

// TODO: testing
func (r *UserController) Index(ctx http.Context) http.Response {
	page := ctx.Request().QueryInt("page", 1)
	page_size := ctx.Request().QueryInt("page_size", 10)
	var users []models.User
	var total int64

	facades.Orm().Query().Paginate(page, page_size, &users, &total)

	var last_page int64
	if total%int64(page_size) == 0 {
		last_page = total / int64(page_size)
	} else {
		last_page = total/int64(page_size) + 1
	}

	return ctx.Response().Success().Json(http.Json{
		"message": "Users found successfully.",
		"data":    users,
		"meta": http.Json{
			"total":        total,
			"per_page":     page_size,
			"current_page": page,
			"last_page":    last_page,
		},
	})
}

func (r *UserController) Show(ctx http.Context) http.Response {
	var id = ctx.Request().Route("id")

	var user models.User
	// facades.Orm().Query().Find(&user, id)
	facades.Orm().Query().Where("id = ?", id).First(&user)

	if user.ID == 0 {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "User not found.",
			"data":    nil,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": "User found successfully.",
		"data":    user,
	})
}

func (r *UserController) Store(ctx http.Context) http.Response {
	/*
	* validation request
	 */
	validator, _ := facades.Validation().Make(
		map[string]any{
			"full_name": ctx.Request().Input("full_name"),
			"email":     ctx.Request().Input("email"),
			"password":  ctx.Request().Input("password"),
			"role":      ctx.Request().Input("role"),
			"is_active": ctx.Request().InputBool("is_active"),
		},
		map[string]string{
			"full_name": "required|string|max_len:255",
			"email":     "required|email|max_len:255",
			"password":  "string|max_len:255",
			"role":      "string|max_len:255",
			"is_active": "bool",
		},
		validation.Attributes(map[string]string{
			"full_name": "Full Name",
			"email":     "Email Address",
			"password":  "Password",
			"role":      "Role",
			"is_active": "Is Active",
		}),
	)

	if validator.Fails() {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "User create failed.",
			"errors":  validator.Errors().All(),
		})
	}

	/*
	* validation unique email
	 */
	email := ctx.Request().Input("email")
	if !helpers.UniqueEmail(email) {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "User create failed.",
			"errors":  "The email has already been taken.",
		})
	}

	/*
	* bind user
	 */
	var user models.User
	err := validator.Bind(&user)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "User create failed when bind data.",
			"data":    nil,
			"err":     err,
		})
	}

	/*
	* create user
	 */
	err = facades.Orm().Query().Create(&user)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "User create failed.",
			"data":    nil,
			"err":     err,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": "User created successfully.",
		"data":    user,
		// TODO: add meta response about pagination
	})
}

func (r *UserController) Update(ctx http.Context) http.Response {
	/*
	* validation request
	 */
	validator, _ := facades.Validation().Make(
		map[string]any{
			"full_name": ctx.Request().Input("full_name"),
			"email":     ctx.Request().Input("email"),
			"password":  ctx.Request().Input("password"),
			"role":      ctx.Request().Input("role"),
			"is_active": ctx.Request().InputBool("is_active"),
		},
		map[string]string{
			"full_name": "string|max_len:255",
			"email":     "email|max_len:255",
			"password":  "string|max_len:255",
			"role":      "string|max_len:255",
			"is_active": "bool",
		},
		validation.Attributes(map[string]string{
			"full_name": "Full Name",
			"email":     "Email Address",
			"password":  "Password",
			"role":      "Role",
			"is_active": "Is Active",
		}),
	)

	if validator.Fails() {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "User update failed.",
			"errors":  validator.Errors().All(),
		})
	}

	/*
	* validation id params
	 */
	var id = ctx.Request().Route("id")
	var user models.User

	facades.Orm().Query().Where("id = ?", id).First(&user)

	if user.ID == 0 {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "User not found.",
			"data":    nil,
		})
	}

	/*
	* validation unique email
	 */
	email := ctx.Request().Input("email")
	if email != user.Email && email != "" && !helpers.UniqueEmail(email) {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "User create failed.",
			"errors":  "The email has already been taken.",
		})
	}

	/*
	* bind user
	 */
	err := validator.Bind(&user)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "User create failed.",
			"data":    nil,
			"err":     err,
		})
	}

	/*
	* update user
	 */
	_, err = facades.Orm().Query().Where("id = ?", id).Update(&user)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "User update failed.",
			"data":    nil,
		})
	}

	/*
	* get user and return response
	 */
	facades.Orm().Query().Where("id = ?", id).First(&user)

	return ctx.Response().Success().Json(http.Json{
		"message": "User updated successfully.",
		"data":    user,
	})
}

func (r *UserController) Destroy(ctx http.Context) http.Response {
	var id = ctx.Request().Route("id")

	var user models.User
	res, err := facades.Orm().Query().Where("id = ?", id).Delete(&user)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "User delete failed.",
			"data":    nil,
		})
	}

	if res.RowsAffected == 0 {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "User not found.",
			"data":    nil,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": "User deleted successfully.",
		"data":    nil,
	})
}
