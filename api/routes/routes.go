package routes

import (
	"github.com/labstack/echo"
	"github.com/rof20004/brasil-servicos/api/modules/user"
)

// API program
var API = echo.New()

// InitRoutes start
func InitRoutes() {
	// Users routes
	API.GET("users", user.ListUsers)
	API.POST("users", user.InsertUser)
	API.DELETE("users/:id", user.DeleteUser)
}
