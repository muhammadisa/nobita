package feature

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/feature"
	"github.com/muhammadisa/nobita/util/dbtrx"
)

func (r *rw) ReadFeature(ctx context.Context, roleID int64, name string)  error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer dbtrx.Trx(tx, err)

	var featr feature.Feature
	const query = `SELECT * FROM features WHERE role_id = ? AND name = ?`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	mutex.Lock()
	rows, err := stmt.QueryContext(ctx, roleID, name)
	if err != nil {
		return err
	}
	mutex.Unlock()
	for rows.Next() {
		err = rows.Scan(featr.FastScan()...)
		if err != nil {
			return err
		}
	}
	return nil
}

