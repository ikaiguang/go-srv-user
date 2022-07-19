// Package repos
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package repos

import (
	context "context"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"

	entities "github.com/ikaiguang/go-srv-user/internal/domain/entity/org"
)

// OrgRepo repo
type OrgRepo interface {
	Create(ctx context.Context, dataModel *entities.Org) error
	ExistCreate(ctx context.Context, dataModel *entities.Org) (anotherModel *entities.Org, isNotFound bool, err error)
	CreateInBatches(ctx context.Context, dataModels []*entities.Org, batchSize int) error
	Insert(ctx context.Context, dataModels []*entities.Org) error
	Update(ctx context.Context, dataModel *entities.Org) error
	ExistUpdate(ctx context.Context, dataModel *entities.Org) (anotherModel *entities.Org, isNotFound bool, err error)
	QueryOneById(ctx context.Context, id interface{}) (dataModel *entities.Org, isNotFound bool, err error)
	QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *entities.Org, isNotFound bool, err error)
	QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) (dataModels []*entities.Org, err error)
	List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) (dataModels []*entities.Org, totalNumber int64, err error)
	Delete(ctx context.Context, dataModel *entities.Org) error
	DeleteByIds(ctx context.Context, ids interface{}) error
}
