package routes

import (
	"backend-go/src/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ItemsRoutes(e *echo.Echo, db *gorm.DB) {
	items := e.Group("/items")

	items.GET("", func(c echo.Context) error {
		return controller.GetAllItems(c, db)
	})

	items.POST("", func(c echo.Context) error {
		return controller.AddItem(c, db)
	})

	items.GET("/filter", func(c echo.Context) error {
		return controller.FilterCategory(c, db)
	})

	items.DELETE("/:id", controller.DeleteItem)

	items.PATCH("/:id", func(c echo.Context) error {
		return controller.UpdateItem(c, db)
	})

	// items.DELETE("/:id", func(c echo.Context) error {
	// 	return controller.DeleteItem(c, db)
	// })

}
