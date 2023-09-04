package transactionservice

import "github.com/google/uuid"

type TransactionReqUpdate struct {
	ID uuid.UUID `json:"id"`
}
