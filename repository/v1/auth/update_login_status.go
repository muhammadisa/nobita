package auth

import (
	"context"
	"github.com/muhammadisa/nobita/util/dbtrx"
)

func (r *rw) UpdateLoginStatus(ctx context.Context, accountID int64, status bool) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer dbtrx.Trx(tx, err)

	var query = `UPDATE accounts SET logged_in = ? WHERE id = ?`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.ExecContext(ctx, status, accountID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil && rowsAffected == 0 {
		return err
	}
	return nil
}
