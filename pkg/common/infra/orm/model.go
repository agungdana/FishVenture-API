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
