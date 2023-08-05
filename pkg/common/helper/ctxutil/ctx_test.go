package ctxutil_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func Test_Context(t *testing.T) {
	ctx := context.Background()

	userID := uuid.New()
	roleID_1 := uuid.New()
	roleID_2 := uuid.New()

	ctx = ctxutil.NewRequestID(ctx)
	ctx = ctxutil.NewTransactionID(ctx)
	ctx = ctxutil.SetUserID(ctx, userID)
	ctx = ctxutil.SetRoleID(ctx, roleID_1, roleID_2)

	idUserFromCtx, ok := ctxutil.GetUserID(ctx)

	assert.Equal(t, userID, idUserFromCtx, fmt.Sprintf("status : %v", ok))

	idRoleFromCtx, ok := ctxutil.GetRoleID(ctx)

	assert.Equal(t, []uuid.UUID{roleID_1, roleID_2}, idRoleFromCtx, fmt.Sprintf("status : %v", ok))

	idRequestFromCtx, ok := ctxutil.GetRequestID(ctx)

	assert.True(t, ok, "requestID %v: - status %v:", idRequestFromCtx, ok)

	idTransactionFromCtx, ok := ctxutil.GetTransactionID(ctx)

	assert.True(t, ok, "transactionID %v: - status %v:", idTransactionFromCtx, ok)

}

func Benchmark_ctxutil(b *testing.B) {
	ctx := context.Background()

	userID := uuid.New()
	pondID := uuid.New()
	roleID := []uuid.UUID{
		uuid.New(), uuid.New(),
	}

	ctx = ctxutil.NewRequest(ctx)
	ctx = ctxutil.SetUserPayload(ctx, userID, pondID, "", roleID...)

	idUserFromCtx, ok := ctxutil.GetUserID(ctx)

	assert.Equal(b, userID, idUserFromCtx, fmt.Sprintf("status : %v", ok))

	idRoleFromCtx, ok := ctxutil.GetRoleID(ctx)

	assert.Equal(b, roleID, idRoleFromCtx, fmt.Sprintf("status : %v", ok))

	idRequestFromCtx, ok := ctxutil.GetRequestID(ctx)

	assert.True(b, ok, "requestID %v: - status %v:", idRequestFromCtx, ok)

	idTransactionFromCtx, ok := ctxutil.GetTransactionID(ctx)

	assert.True(b, ok, "transactionID %v: - status %v:", idTransactionFromCtx, ok)

}
