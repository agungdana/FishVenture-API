package ctxutil_test

import (
	"context"
	"testing"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

type Role struct {
	ID             uuid.UUID
	Number         string
	Name           string
	Scope          string
	RolePermission []RolePermission
}

type Permission struct {
	ID             uuid.UUID
	Number         string
	Name           string
	RolePermission []RolePermission
}

type RolePermission struct {
	ID               uuid.UUID
	RoleID           uuid.UUID
	Role             Role
	PermissionID     uuid.UUID
	PermissionNumber string
	PermissionName   string
	Permission       Permission
}

type UserPermission struct {
	ID               uuid.UUID
	UserID           uuid.UUID
	PermissionID     uuid.UUID
	PermissionNumber string
	PermissionName   string
	Permission       Permission
}

var (
	userID_1      = uuid.New()
	userID_2      = uuid.New()
	roleID_1      = uuid.New()
	roleID_2      = uuid.New()
	PermisionID_1 = uuid.New()
	PermisionID_2 = uuid.New()
	PermisionID_3 = uuid.New()
	PermisionID_4 = uuid.New()
	PermisionID_5 = uuid.New()
)

var ListOfRoles = []Role{
	{
		ID:     roleID_1,
		Number: "R0001",
		Name:   "Admin",
		Scope:  "GLOBAL",
	}, {
		ID:     roleID_2,
		Number: "R0002",
		Name:   "Customer",
		Scope:  "GLOBAL",
	},
}

var ListOfPermision = []Permission{
	{
		ID:     PermisionID_1,
		Number: "P0001",
		Name:   "/create",
	}, {
		ID:     PermisionID_2,
		Number: "P0002",
		Name:   "/delete",
	}, {
		ID:     PermisionID_3,
		Number: "P0003",
		Name:   "/update",
	}, {
		ID:     PermisionID_4,
		Number: "P0004",
		Name:   "/read",
	}, {
		ID:     PermisionID_5,
		Number: "P0005",
		Name:   "/read-all",
	},
}

var ListOfRolePermision = []RolePermission{
	{
		ID:               uuid.New(),
		RoleID:           roleID_1,
		PermissionID:     PermisionID_1,
		PermissionNumber: "P0001",
		PermissionName:   "/create",
		Permission:       Permission{},
	}, {
		ID:               uuid.New(),
		RoleID:           roleID_1,
		PermissionID:     PermisionID_2,
		PermissionNumber: "P0002",
		PermissionName:   "/delete",
		Permission:       Permission{},
	}, {
		ID:               uuid.New(),
		RoleID:           roleID_1,
		PermissionID:     PermisionID_3,
		PermissionNumber: "P0003",
		PermissionName:   "/update",
		Permission:       Permission{},
	}, {
		ID:               uuid.New(),
		RoleID:           roleID_1,
		PermissionID:     PermisionID_4,
		PermissionNumber: "P0004",
		PermissionName:   "/read",
		Permission:       Permission{},
	}, {
		ID:               uuid.New(),
		RoleID:           roleID_1,
		PermissionID:     PermisionID_5,
		PermissionNumber: "P0005",
		PermissionName:   "/read-all",
		Permission:       Permission{},
	}, {
		ID:               uuid.New(),
		RoleID:           roleID_2,
		PermissionID:     PermisionID_4,
		PermissionNumber: "P0005",
		PermissionName:   "/read-all",
		Permission:       Permission{},
	}, {
		ID:               uuid.New(),
		RoleID:           roleID_2,
		PermissionID:     PermisionID_5,
		PermissionNumber: "P0005",
		PermissionName:   "/read-all",
		Permission:       Permission{},
	},
}

var listOfUserPermission = []UserPermission{
	// {
	// 	ID:               uuid.New(),
	// 	UserID:           userID_1,
	// 	PermissionID:     PermisionID_1,
	// 	PermissionNumber: "P0001",
	// 	PermissionName:   "/create",
	// },
	{
		ID:               uuid.New(),
		UserID:           userID_2,
		PermissionID:     PermisionID_2,
		PermissionNumber: "P0001",
		PermissionName:   "/delete",
	},
}

func Test_ValidateAccessFalse(t *testing.T) {

	permissionAccess := []ctxutil.PermissionAccess{}

	for _, v := range ListOfRolePermision {
		permissionAccess = append(permissionAccess, ctxutil.PermissionAccess{
			ID:   v.RoleID,
			Path: v.PermissionName,
		})
	}
	ctxutil.AddPermissionAccess(permissionAccess)

	ctx := context.Background()
	userID := uuid.New()
	appType := ""

	ctx = ctxutil.SetUserPayload(ctx, userID, roleID_2, appType)

	ok := ctxutil.CanAccess(ctx, "/update")

	assert.False(t, ok, "value: %v", ok)

}

func Test_ValidateAccessTrue(t *testing.T) {

	permissionAccess := []ctxutil.PermissionAccess{}

	for _, v := range ListOfRolePermision {
		permissionAccess = append(permissionAccess, ctxutil.PermissionAccess{
			ID:   v.RoleID,
			Path: v.PermissionName,
		})
	}
	ctxutil.AddPermissionAccess(permissionAccess)

	ctx := context.Background()
	userID := uuid.New()
	appType := ""

	ctx = ctxutil.SetUserPayload(ctx, userID, roleID_1, appType)

	ok := ctxutil.CanAccess(ctx, "/update")

	assert.True(t, ok, "value: %v", ok)

}

func Test_UserPermission(t *testing.T) {

	permissionAccess := []ctxutil.PermissionAccess{}

	for _, v := range ListOfRolePermision {
		permissionAccess = append(permissionAccess, ctxutil.PermissionAccess{
			ID:   v.RoleID,
			Path: v.PermissionName,
		})
	}

	for _, v := range listOfUserPermission {
		permissionAccess = append(permissionAccess, ctxutil.PermissionAccess{
			ID:   v.UserID,
			Path: v.PermissionName,
		})
	}
	ctxutil.AddPermissionAccess(permissionAccess)

	ctx := context.Background()

	ctx = ctxutil.SetUserPayload(ctx, userID_1, roleID_2, "")

	ok := ctxutil.CanAccess(ctx, "/create")

	assert.True(t, ok, "value: %v", ok)

}

func Benchmark_UserPermission(b *testing.B) {

	permissionAccess := []ctxutil.PermissionAccess{}

	for _, v := range ListOfRolePermision {
		permissionAccess = append(permissionAccess, ctxutil.PermissionAccess{
			ID:   v.RoleID,
			Path: v.PermissionName,
		})
	}

	for _, v := range listOfUserPermission {
		permissionAccess = append(permissionAccess, ctxutil.PermissionAccess{
			ID:   v.UserID,
			Path: v.PermissionName,
		})
	}
	ctxutil.AddPermissionAccess(permissionAccess)

	ctx := context.Background()

	ctx = ctxutil.SetUserPayload(ctx, userID_1, roleID_2, "")

	ok := ctxutil.CanAccess(ctx, "/create")

	assert.True(b, ok, "value: %v", ok)

}

func Test_DeleteUserPermission(t *testing.T) {

	permissionAccess := []ctxutil.PermissionAccess{}

	for _, v := range ListOfRolePermision {
		permissionAccess = append(permissionAccess, ctxutil.PermissionAccess{
			ID:   v.RoleID,
			Path: v.PermissionName,
		})
	}

	for _, v := range listOfUserPermission {
		permissionAccess = append(permissionAccess, ctxutil.PermissionAccess{
			ID:   v.UserID,
			Path: v.PermissionName,
		})
	}

	ctxutil.AddPermissionAccess(permissionAccess)

	ctx := context.Background()

	ctx = ctxutil.SetUserPayload(ctx, userID_1, roleID_2, "")

	ok := ctxutil.CanAccess(ctx, "/create")

	assert.True(t, ok, "value: %v", ok)

	ctxutil.DeleteUserPermission(ctxutil.PermissionAccess{
		ID:   userID_1,
		Path: "/create",
	})

	ok = ctxutil.CanAccess(ctx, "/create")

	assert.True(t, ok, "value after delete: %v", ok)

}

func Test_DeleteUserRole(t *testing.T) {

	permissionAccess := []ctxutil.PermissionAccess{}

	for _, v := range ListOfRolePermision {
		permissionAccess = append(permissionAccess, ctxutil.PermissionAccess{
			ID:   v.RoleID,
			Path: v.PermissionName,
		})
	}

	for _, v := range listOfUserPermission {
		permissionAccess = append(permissionAccess, ctxutil.PermissionAccess{
			ID:   v.UserID,
			Path: v.PermissionName,
		})
	}

	ctxutil.AddPermissionAccess(permissionAccess)

	ctx := context.Background()

	ctx = ctxutil.SetUserPayload(ctx, userID_1, roleID_1, "")

	ok := ctxutil.CanAccess(ctx, "/create")

	assert.True(t, ok, "value: %v", ok)

	ctxutil.DeleteRolePermission(ctxutil.PermissionAccess{
		ID: roleID_1,
	})

	ok = ctxutil.CanAccess(ctx, "/create")

	assert.True(t, ok, "value after delete: %v", ok)

}
