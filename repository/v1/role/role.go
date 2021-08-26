package role

import (
	"database/sql"
	rolerwinterface "github.com/muhammadisa/nobita/repository/v1/role/interface"
	"github.com/muhammadisa/nobita/util/dbc"
	"sync"
)

var mutex = &sync.RWMutex{}

type rw struct {
	db *sql.DB
}

func NewRoleRepo(config dbc.Config) (rolerwinterface.RW, error) {
	db, err := dbc.OpenDB(config)
	if err != nil {
		return nil, err
	}
	return &rw{
		db: db,
	}, nil
}
