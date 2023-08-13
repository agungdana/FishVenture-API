package model

import (
	"fmt"
	"time"

	"github.com/e-fish/api/pkg/common/helper/rand"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	errortransaction "github.com/e-fish/api/pkg/domain/transaction/error-transaction"
	"github.com/google/uuid"
)

type CreateOrderInput struct {
	BudidayaID  uuid.UUID  `json:"budidayaID"`
	Qty         int        `json:"qty"`
	BookingDate *time.Time `json:"bookingDate"`
}

func (c *CreateOrderInput) Validate() error {
	errs := werror.NewError("failed validate error")

	if c.BudidayaID == uuid.Nil {
		errs.Add(errortransaction.ErrValidateCreateInput.AttacthDetail(map[string]any{"budidayaID": "empty"}))
	}
	if c.Qty == 0 {
		errs.Add(errortransaction.ErrValidateCreateInput.AttacthDetail(map[string]any{"qty": "empty"}))
	}
	if c.BookingDate == nil {
		errs.Add(errortransaction.ErrValidateCreateInput.AttacthDetail(map[string]any{"bookingDate": "empty"}))
	}

	return errs.Return()
}

func (c *CreateOrderInput) ToOrder(userID uuid.UUID, pricelist model.PriceList) Order {
	return Order{
		ID:          uuid.New(),
		Code:        GenerateCode(),
		PondID:      pricelist.Budidaya.PondID,
		BudidayaID:  c.BudidayaID,
		UserID:      userID,
		Qty:         c.Qty,
		PricelistID: pricelist.ID,
		Price:       float64(pricelist.Price),
		Ammout:      float64(pricelist.Price) * float64(c.Qty),
		BookingDate: c.BookingDate,
		Status:      ACTIVE,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

func GenerateCode() string {
	today := time.Now()
	year := today.Year()
	month := today.Month()
	day := today.Weekday().String()
	newDay := day[0:3]
	return fmt.Sprintf("%v-%v%d/%v/%v", "OC", year, month, newDay, rand.RandCode(4))
}

type ReadInput struct {
	orm.Paginantion
}
