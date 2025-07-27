package models

import "time"

type Item struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	CategoryItem string     `json:"category_item"`  // 'tools' | 'materials'
	ItemName     string     `json:"item_name"`      // Nama barang
	ItemPicture  *string    `json:"item_picture"`   // Path atau URL gambar (nullable)
	Quantity     int        `json:"quantity"`       // Jumlah stok
	Unit         *string    `json:"unit"`           // Satuan (nullable)
	Price        *float64   `json:"price"`          // Harga satuan (nullable)
	GoodsInBy    uint       `json:"goods_in_by"`    // Nama user yang input barang masuk
	GoodsOutBy   *uint      `json:"goods_out_by"`   // Nama user yang input barang keluar
	TimeGoodsIn  time.Time  `json:"time_goods_in"`  // Timestamp barang masuk
	TimeGoodsOut *time.Time `json:"time_goods_out"` // Nullable, saat barang keluar
	CreatedAt    time.Time  `json:"created_at"`     // Timestamp buat data
	UpdatedAt    time.Time  `json:"updated_at"`     // Timestamp update data

	GoodsInUser  *User `gorm:"foreignKey:GoodsInBy"`
	GoodsOutUser *User `gorm:"foreignKey:GoodsOutBy"`
}

func (Item) TableName() string {
	return "inventory_management.items"
}
