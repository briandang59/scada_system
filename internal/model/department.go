package model

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID        uint           `json:"id"`
	NameEn    string         `json:"name_en" `
	NameZh    string         `json:"name_zh" `
	NameVn    string         `json:"name_vn" `
	Code      string         `json:"code"`
	Active    bool           `json:"active" gorm:"default:true"`
	FactoryID uint           `json:"factory_id"`
	Factory   *Factory       `json:"factory,omitempty" gorm:"foreignKey:FactoryID"`
	ManagerID uint           `json:"manager_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"           gorm:"index"`
}
