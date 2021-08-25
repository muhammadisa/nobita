package auth

import (
	"context"
	"fmt"
	"github.com/muhammadisa/nobita/model/v1/auth"
	"github.com/muhammadisa/nobita/util/dbtrx"
	"github.com/muhammadisa/nobita/util/hash"
)

func (r *rw) VerifyTempSecret(ctx context.Context, secret auth.Secret) (auth.Account, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return auth.Account{}, err
	}
	defer dbtrx.Trx(tx, err)

	// logic
	var k, s string
	var roleID, accountID int64
	var query = `SELECT id, role_id, %s, secrets FROM accounts WHERE %s = ?`
	stmt, err := tx.Prepare(fmt.Sprintf(query, secret.Kind, secret.Kind))
	if err != nil {
		return auth.Account{}, err
	}
	rows := stmt.QueryRowContext(ctx, secret.Identifier)
	err = rows.Scan(&accountID, &roleID, &k, &s)
	if err != nil {
		return auth.Account{}, err
	}
	err = hash.Verify(s, secret.Code)
	if err != nil {
		return auth.Account{}, err
	}
	return auth.Account{
		ID:     accountID,
		RoleID: roleID,
	}, nil
}
