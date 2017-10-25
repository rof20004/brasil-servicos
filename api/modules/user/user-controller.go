package user

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// ListUsers - list all users
func ListUsers(c echo.Context) error {
	users, err := List()
	if err != nil {
		log.Println("[ListUsers]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Ocorreu um erro ao realizar consulta")
	}
	return c.JSON(http.StatusOK, users)
}

// InsertUser - create a user
func InsertUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		log.Println("[InsertUser]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Informações incompletas ou inválidas")
	}

	err := Insert(user)
	if err != nil {
		log.Println("[InsertUser]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Ocorreu um erro ao realizar o cadastro")
	}

	return c.String(http.StatusOK, "Usuário cadastrado com sucesso")
}

// DeleteUser - delete a user
func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := Delete(id)
	if err != nil {
		log.Println("[DeleteUser]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Ocorreu um erro ao remover usuário")
	}

	return c.String(http.StatusOK, "Usuário removido com sucesso")
}
