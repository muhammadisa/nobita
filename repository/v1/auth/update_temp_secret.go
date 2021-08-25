package auth

import (
	"context"
	"github.com/muhammadisa/nobita/util/dbtrx"
	"time"
)

func (r *rw) UpdateTempSecret(ctx context.Context, accountID int64, hashedSecret string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer dbtrx.Trx(tx, err)

	var query = `UPDATE accounts SET secrets = ?, updated_at = ? WHERE id = ?`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.ExecContext(ctx, hashedSecret, time.Now(), accountID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil && rowsAffected == 0 {
		return err
	}
	return nil
}
