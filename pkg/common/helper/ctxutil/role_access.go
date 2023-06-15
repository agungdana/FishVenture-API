package ctxutil

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

var (
	permissionAccess = make(map[uuid.UUID]map[string]bool)

	mut sync.Mutex
)

type PermissionAccess struct {
	ID   uuid.UUID
	Path string
}

func DeleteUserPermission(data PermissionAccess) {
	mut.Lock()
	if val, ok := permissionAccess[data.ID]; ok {
		delete(val, data.Path)
	}
	mut.Unlock()
}

func DeleteRolePermission(data PermissionAccess) {
	mut.Lock()
	delete(permissionAccess, data.ID)
	mut.Unlock()
}

func AddPermissionAccess(data []PermissionAccess) {
	mut.Lock()

	for _, v := range data {
		if _, ok := permissionAccess[v.ID]; !ok {
			permissionAccess[v.ID] = make(map[string]bool)
		}

		if _, ok := permissionAccess[v.ID][v.Path]; !ok {
			permissionAccess[v.ID][v.Path] = true
		}
	}

	mut.Unlock()
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
