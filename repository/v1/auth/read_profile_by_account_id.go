package auth

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/auth"
)

func (r *rw) ReadProfileByAccountID(ctx context.Context, accountID int64) (auth.Profile, error) {
	const query = `SELECT * FROM profilers WHERE account_id = ?`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return auth.Profile{}, err
	}
	mutex.Lock()
	rows, err := stmt.QueryContext(ctx, accountID)
	if err != nil {
		return auth.Profile{}, err
	}
	mutex.Unlock()
	var profile auth.Profile
	for rows.Next() {
		err = rows.Scan(profile.FastScan()...)
		if err != nil {
			return auth.Profile{}, err
		}
		profile.UseUnixTimestamp()
	}
	return profile, nil
}
