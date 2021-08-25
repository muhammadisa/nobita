package auth

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/auth"
	"github.com/muhammadisa/nobita/util/dbtrx"
	"time"
)

func (r *rw) UpdateProfileByAccountID(ctx context.Context, profile auth.Profile) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer dbtrx.Trx(tx, err)

	var query = `
		UPDATE profilers SET 
			long_lat = ?,
			full_name = ?,
			birth_day = ?,
			gender = ?,
			updated_at =?
		WHERE account_id = ?
	`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err

	}
	result, err := stmt.ExecContext(ctx, profile.LongLat, profile.FullName, profile.BirthDay, profile.Gender, time.Now(), profile.AccountID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil && rowsAffected == 0 {
		return err
	}
	return nil
}
