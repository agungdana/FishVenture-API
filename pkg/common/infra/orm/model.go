package orm

import (
	"time"

	"github.com/google/uuid"
)

type OrmModel struct {
	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt *time.Time
	UpdatedBy *uuid.UUID
	DeletedAt *time.Time `sql:"index"`
	DeletedBy *uuid.UUID
}

type Paginantion struct {
	FindBy      string `json:"findBy"`
	Keyword     string `json:"keyword"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	Sort        string `json:"sort"`
	Direction   string `json:"direction"`
	TotalRows   int64  `json:"totalRows"`
	TotalPage   int    `json:"totalPage"`
	ObjectTable any    `json:"objectTable"`
}
