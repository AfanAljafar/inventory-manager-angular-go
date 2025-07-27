package main

import (
	"backend-go/src/config"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"backend-go/src/middlewareCustom"
	"backend-go/src/routes"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":4001"
	}

	config.DataBaseConnection()

	db := config.GetDB()

	e.Use(middleware.CORS()) // Allow all
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Custom Logging Middleware
	e.Use(middlewareCustom.EventLogging)

	// Routes
	routes.UsersRoutes(e, db)
	routes.ItemsRoutes(e, db)

	// Route GET /
	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, Message{Message: "Hello from Echo GET!"})
	// })

	fmt.Printf("✅ Example app listening on port %s\n", PORT)

	e.Logger.Fatal(e.Start(PORT))
}
