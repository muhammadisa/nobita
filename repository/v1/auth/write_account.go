package auth

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/auth"
	"github.com/muhammadisa/nobita/util/dbtrx"
)

func (r *rw) WriteAccount(ctx context.Context, request auth.Account) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer dbtrx.Trx(tx, err)

	var query = `INSERT INTO accounts(email, phone, role_id, device_id) VALUES (?,?,?,?)`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return 0, err
	}
	result, err := stmt.ExecContext(ctx, request.Email, request.Phone, request.RoleID, request.DeviceID)
	if err != nil {
		return 0, err
	}
	insertedID, err := result.LastInsertId()
	if err != nil && insertedID == 0 {
		return 0, err
	}
	return insertedID, nil
}
