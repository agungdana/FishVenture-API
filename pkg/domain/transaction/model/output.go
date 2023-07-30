package model

import (
	"time"

	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/google/uuid"
)

type OrderOutput struct {
	ID          uuid.UUID             `gorm:"primaryKey,size:256" json:"id"`
	Code        string                `json:"code"`
	BudidayaID  uuid.UUID             `json:"budidayaID"`
	Budidaya    model.BudidayaOutput  `gorm:"foreignKey:BudidayaID;references:ID" json:"budidaya"`
	User        User                  `json:"user"`
	Qty         int                   `json:"qty"`
	BookingDate *time.Time            `json:"bookingDate"`
	PricelistID uuid.UUID             `json:"pricelistID"`
	Pricelist   model.PriceListOutput `gorm:"foreignKey:PricelistID;references:ID" json:"pricelist"`
	Price       float64               `json:"price"`
	Ammout      float64               `json:"ammout"`
	Status      string                `json:"status"`
	CreatedAt   time.Time             `json:"createdAt"`
}

func (*OrderOutput) TableName() string {
	return "order"
}

type OrderOutputPagination struct {
	FindBy    string `json:"findBy"`
	Keyword   string `json:"keyword"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"`
	TotalRows int64  `json:"totalRows"`
	TotalPage int    `json:"totalPage"`
	Rows      any    `json:"rows"`
}
