package controller

import (
	"backend-go/src/config"
	"backend-go/src/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllItems(c echo.Context, db *gorm.DB) error {
	items, err := repository.GetAllItems(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch items",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success fetching items",
		"data":    items,
	})
}

func AddItem(c echo.Context, db *gorm.DB) error {
	item, err := repository.AddItem(c, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":  "Failed to add item",
			"detail": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Item added successfully",
		"data":    item,
	})
}

func FilterCategory(c echo.Context, db *gorm.DB) error {
	category := c.QueryParam("category")
	if category == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Category query parameter is required",
		})
	}

	items, err := repository.FilterItemsByCategory(db, category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to filter items by category",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success filtering items by category",
		"data":    items,
	})
}

func DeleteItem(c echo.Context) error {
	db := config.GetDB()

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Item ID is required",
		})
	}

	err := repository.DeleteItem(db, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":  "Failed to delete item",
			"detail": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Item deleted successfully",
	})
}

func UpdateItem(c echo.Context, db *gorm.DB) error {
	item, err := repository.UpdateItem(c, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":  "Failed to update item",
			"detail": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Item updated successfully",
		"data":    item,
	})
}
