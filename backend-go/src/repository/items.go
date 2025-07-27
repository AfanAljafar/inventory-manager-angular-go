package repository

import (
	"backend-go/src/models"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllItems(db *gorm.DB) ([]models.Item, error) {
	var items []models.Item
	if err := db.Preload("GoodsInUser").Preload("GoodsOutUser").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func AddItem(c echo.Context, db *gorm.DB) (*models.Item, error) {
	var item models.Item
	if err := c.Bind(&item); err != nil {
		return nil, err
	}

	// Atur waktu masuk otomatis jika kosong
	if item.TimeGoodsIn.IsZero() {
		item.TimeGoodsIn = time.Now()
	}

	// Optional: jika GoodsOutBy == 0, anggap null
	if item.GoodsOutBy != nil && *item.GoodsOutBy == 0 {
		item.GoodsOutBy = nil
	}

	if err := db.Create(&item).Error; err != nil {
		return nil, err
	}

	// Re-fetch dengan preload relasi user (biar langsung ada relasi user di response)
	if err := db.
		Preload("GoodsInUser").
		Preload("GoodsOutUser").
		First(&item, item.ID).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func FilterItemsByCategory(db *gorm.DB, category string) ([]models.Item, error) {
	var items []models.Item
	if err := db.Preload("GoodsInUser").Preload("GoodsOutUser").
		Where("LOWER(category_item) = LOWER(?)", category).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func DeleteItem(db *gorm.DB, id string) error {
	if id == "" {
		return echo.NewHTTPError(400, "Item ID is required")
	}
	// Pastikan item ada sebelum dihapus
	var item models.Item
	if err := db.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(404, "Item not found")
		}
	}

	// Hapus item
	if err := db.Where("id = ?", id).Delete(&models.Item{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateItem(c echo.Context, db *gorm.DB) (*models.Item, error) {
	var item models.Item
	if err := c.Bind(&item); err != nil {
		return nil, echo.NewHTTPError(400, "Invalid request body")
	}

	// Atur waktu masuk otomatis jika kosong
	if item.TimeGoodsIn.IsZero() {
		item.TimeGoodsIn = time.Now()
	}

	// Optional: jika GoodsOutBy == 0, anggap null
	if item.GoodsOutBy != nil && *item.GoodsOutBy == 0 {
		item.GoodsOutBy = nil
	}

	if item.ID == 0 {
		return nil, echo.NewHTTPError(400, "Item ID is required for update")
	}

	// Update item in the database
	if err := db.Save(&item).Error; err != nil {
		return nil, echo.NewHTTPError(500, "Failed to update item")
	}

	// Re-fetch with preload to return updated item with relations
	if err := db.Preload("GoodsInUser").Preload("GoodsOutUser").First(&item, item.ID).Error; err != nil {
		return nil, echo.NewHTTPError(500, "Failed to fetch updated item")
	}

	return &item, nil
}
