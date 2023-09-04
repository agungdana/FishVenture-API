package model

import (
	"github.com/google/uuid"
)

type BannerOutput struct {
	ID          uuid.UUID `gorm:"primaryKey,size:256" json:"id"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
}

func (c *BannerOutput) TableName() string {
	return "banners"
}
