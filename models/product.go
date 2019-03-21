package models

import (
	"github.com/jinzhu/gorm"
)

type ProductModel struct {
	gorm.Model
	Brand   string `gorm:"column:brand" json:"brand"`
	Image   string `gorm:"column:image" json:"image"`
	ItemSKU string `gorm:"column:item_sku" json:"itemSku"`
	Name    string `gorm:"column:name" json:"name"`
}

func (ProductModel) TableName() string {
	return "products"
}
