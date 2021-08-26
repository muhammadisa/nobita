package feature

import (
	"database/sql"
	featurerwinterface "github.com/muhammadisa/nobita/repository/v1/feature/interface"
	"github.com/muhammadisa/nobita/util/dbc"
	"sync"
)

var mutex = &sync.RWMutex{}

type rw struct {
	db *sql.DB
}

func NewFeatureRepo(config dbc.Config) (featurerwinterface.RW, error) {
	db, err := dbc.OpenDB(config)
	if err != nil {
		return nil, err
	}
	return &rw{
		db: db,
	}, nil
}
