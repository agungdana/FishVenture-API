package ctxutil

import (
	"context"

	"github.com/google/uuid"
)

var permissionAccess = make(map[uuid.UUID]map[string]bool)

type PermissionAccess struct {
	ID   uuid.UUID
	Path string
}

func AddPermissionAccess(data []PermissionAccess) {
	for _, v := range data {
		if _, ok := permissionAccess[v.ID]; !ok {
			permissionAccess[v.ID] = make(map[string]bool)
		}

		if _, ok := permissionAccess[v.ID][v.Path]; !ok {
			permissionAccess[v.ID][v.Path] = true
		}
	}
}

func CanAccess(ctx context.Context, path string) bool {
	roleID, _ := GetRoleID(ctx)
	userID, _ := GetUserID(ctx)

	if val, ok := permissionAccess[userID][path]; ok {
		return val

	}

	for _, v := range roleID {
		if val, ok := permissionAccess[v][path]; ok {
			return val
		}
	}

	return false
}
