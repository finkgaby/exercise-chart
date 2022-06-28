package repositories

import (
	"context"
	"errors"
	"example/graph/model"
	"example/server/commons"
	"example/server/repositories/db"
)

type Repository interface {
	Save(ctx context.Context, model model.InputDataEntry) (data interface{}, err error)
	Update(ctx context.Context, model interface{}) (data interface{}, err error)
	Get(ctx context.Context) (data []*model.DataEntry, err error)
	GetById(ctx context.Context, id string) (data *model.DataEntry, err error)
}

func GetRepository(repoType commons.RepositoryType) (Repository, error) {
	if repoType == commons.RepositoryTypeDB {
		return db.NewRepository()
	} else {
		return nil, errors.New("Invalid Repository Type provided")
	}
}
