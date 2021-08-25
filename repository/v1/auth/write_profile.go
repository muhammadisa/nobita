package auth

import (
	"context"
	"github.com/muhammadisa/nobita/util/dbtrx"
)

func (r *rw) WriteProfile(ctx context.Context, accountID int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer dbtrx.Trx(tx, err)

	var query = `INSERT INTO profilers(long_lat, account_id, full_name, gender, birth_day) VALUES (?,?,?,?,?)`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.ExecContext(ctx,
		"0:0", accountID, "none",
		"none", "00-00-0000",
	)
	if err != nil {
		return err
	}
	insertedID, err := result.LastInsertId()
	if err != nil && insertedID == 0 {
		return err
	}
	return nil
}
