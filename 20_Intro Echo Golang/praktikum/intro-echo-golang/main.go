package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	userById := GetUserController(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id",
		"user":     userById,
	})
}

func GetUserController(id int) interface{} {
	idKetemu := 0
	for i, value := range users {
		if value.Id == id {
			idKetemu = i
			break
		}
	}

	return users[idKetemu]
}

// delete user by id
func DeleteUserHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	userById := DeleteUserController(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user by id",
		"user":     userById,
	})
}

func DeleteUserController(id int) interface{} {
	idKetemu := 0
	for i, value := range users {
		if value.Id == id {
			idKetemu = i
			break
		}
	}

	users[idKetemu] = users[len(users)-1]
	users[len(users)-1].Id = 0
	users[len(users)-1].Name = ""
	users[len(users)-1].Email = ""
	users[len(users)-1].Password = ""
	users = users[:len(users)-1]

	return users
}

// update user by id
func UpdateUserHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	userById := UpdateUserController(id, c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user by id",
		"user":     userById,
		"users":    users,
	})
}

func UpdateUserController(id int, c echo.Context) interface{} {
	idKetemu := 0
	for i, value := range users {
		if value.Id == id {
			idKetemu = i
			break
		}
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	users[idKetemu].Name = name
	users[idKetemu].Email = email
	users[idKetemu].Password = password

	return users[idKetemu]
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------

func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersHandler)
	e.GET("/users/:id", GetUserHandler)
	e.DELETE("/users/:id", DeleteUserHandler)
	e.PUT("/users/:id", UpdateUserHandler)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
