package auth

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/auth"
	"github.com/muhammadisa/nobita/util/dbtrx"
)

func (r *rw) ReadContactByID(ctx context.Context, accountID int64) (auth.Contact, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return auth.Contact{}, err
	}
	defer dbtrx.Trx(tx, err)

	var contact auth.Contact
	const query = `SELECT id, email, phone FROM accounts WHERE id = ?`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return auth.Contact{}, err
	}
	mutex.Lock()
	rows, err := stmt.QueryContext(ctx, accountID)
	if err != nil {
		return auth.Contact{}, err
	}
	mutex.Unlock()
	for rows.Next() {
		err = rows.Scan(&accountID, &contact.Email, &contact.Phone)
		if err != nil {
			return auth.Contact{}, err
		}
	}
	return contact, nil
}
