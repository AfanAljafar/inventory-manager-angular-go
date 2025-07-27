package routes

import (
	"backend-go/src/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UsersRoutes(e *echo.Echo, db *gorm.DB) {
	users := e.Group("/users")

	users.POST("/login", func(c echo.Context) error {
		return controller.AuthUser(c, db)

	})
}
