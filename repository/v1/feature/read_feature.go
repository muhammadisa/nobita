package feature

import (
	"context"
	"database/sql"
	"errors"
	"github.com/muhammadisa/nobita/model/v1/feature"
)

func (r *rw) ReadFeature(ctx context.Context, roleID int64, name string) error {
	var featr feature.Feature
	const query = `SELECT * FROM features WHERE role_id = ? AND name = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	mutex.Lock()
	row := stmt.QueryRowContext(ctx, roleID, name)
	mutex.Unlock()
	err = row.Scan(featr.FastScan()...)
	if err != nil && err == sql.ErrNoRows {
		return errors.New("this role cannot access this endpoint")
	}
	return nil
}
