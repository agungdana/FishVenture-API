package region

import (
	"context"

	"github.com/e-fish/api/pkg/common/infra/orm"
	errorregion "github.com/e-fish/api/pkg/domain/region/model/error-region"
	"gorm.io/gorm"
)

func newCommand(ctx context.Context, db *gorm.DB) Command {
	var (
		dbTxn = orm.BeginTxn(ctx, db.WithContext(ctx))
	)

	return &command{
		dbTxn: dbTxn,
		query: newQuery(dbTxn),
	}
}

type command struct {
	dbTxn *gorm.DB
	query Query
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return errorregion.ErrCommit.AttacthDetail(map[string]any{"error": err})
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return errorregion.ErrRollback.AttacthDetail(map[string]any{"error": err})
	}
	return nil
}

// CreateNearestTenantByTenantID implements Command.
// func (c *command) CreateNearestTenantByTenantID(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
// 	var (
// 		nearest = []model.CreateNearestTenantInput{}
// 	)

// 	tenant, err := c.tenantQuery.ReadTenantByInputID(ctx, input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	userAddress, err := c.query.ReadAllUserAddressByDistrictID(ctx, tenant.DistrictID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, v := range userAddress {
// 		_, km := distance.Distance(
// 			distance.Coord{
// 				Lat: v.Latitude,
// 				Lon: v.Longitude,
// 			}, distance.Coord{
// 				Lat: tenant.Latitude,
// 				Lon: tenant.Longitude,
// 			},
// 		)

// 		nearest = append(nearest, model.CreateNearestTenantInput{
// 			DistrictID: v.DistrictID,
// 			AddressID:  v.ID,
// 			TenantID:   input,
// 			Distance:   km,
// 		})

// 	}

// 	_, err = c.CreateMultipleNearestTenant(ctx, nearest)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &input, nil
// }

// // CreateNearestTenant implements Command.
// func (c *command) CreateMultipleNearestTenant(ctx context.Context, input []model.CreateNearestTenantInput) ([]uuid.UUID, error) {
// 	var (
// 		userID, _         = ctxutil.GetUserID(ctx)
// 		listNearestTenant = []model.NearestTenant{}
// 		uid               = []uuid.UUID{}
// 	)

// 	for _, v := range input {
// 		newNearestTenant := v.ToNearestTenant(userID)
// 		listNearestTenant = append(listNearestTenant, newNearestTenant)
// 		uid = append(uid, newNearestTenant.ID)
// 	}

// 	err := c.dbTxn.Create(&listNearestTenant).Error
// 	if err != nil {
// 		return nil, errortenant.ErrFailedCreateNearestTenant.AttacthDetail(map[string]any{"error": err})
// 	}

// 	return uid, nil
// }

// func (c *command) CreateNearestTenant(ctx context.Context, input model.NearestTenant) (*uuid.UUID, error) {
// 	err := c.dbTxn.Create(&input).Error
// 	if err != nil {
// 		return nil, err
// 	}
// }
