package auth

import (
	"database/sql"
	authrwinterface "github.com/muhammadisa/nobita/repository/v1/auth/interface"
	"github.com/muhammadisa/nobita/util/dbc"
	"sync"
)

var mutex = &sync.RWMutex{}

type rw struct {
	db *sql.DB
}

func NewAuthRepo(config dbc.Config) (authrwinterface.RW, error) {
	db, err := dbc.OpenDB(config)
	if err != nil {
		return nil, err
	}
	return &rw{
		db: db,
	}, nil
}
