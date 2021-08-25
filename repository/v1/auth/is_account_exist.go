package auth

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/auth"
	"github.com/muhammadisa/nobita/util/dbtrx"
)

func (r *rw) checkEmail(ctx context.Context, query string, request auth.Account) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer dbtrx.Trx(tx, err)

	var id int64
	var emails string
	stmt, err := tx.Prepare(query)
	if err != nil {
		return 0, err
	}
	mutex.Lock()
	rows, err := stmt.QueryContext(ctx, request.Email)
	if err != nil {
		return 0, err
	}
	mutex.Unlock()
	for rows.Next() {
		err = rows.Scan(&id, &emails)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (r *rw) checkPhone(ctx context.Context, query string, request auth.Account) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer dbtrx.Trx(tx, err)

	var id int64
	var phones string
	stmt, err := tx.Prepare(query)
	if err != nil {
		return 0, err
	}
	mutex.Lock()
	rows, err := stmt.QueryContext(ctx, request.Phone)
	if err != nil {
		return 0, err
	}
	mutex.Unlock()
	for rows.Next() {
		err = rows.Scan(&id, &phones)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (r *rw) ReadIsAccountExist(ctx context.Context, request auth.Account) (int64, error) {
	var query string
	if len(request.Email) != 0 {
		query = `SELECT id, email FROM accounts WHERE email = ?`
		return r.checkEmail(ctx, query, request)
	}
	if len(request.Phone) != 0 {
		query = `SELECT id, phone FROM accounts WHERE phone = ?`
		return r.checkPhone(ctx, query, request)
	}
	return 0, nil
}

