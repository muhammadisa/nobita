package role

import "time"

type Role struct {
	ID               int64  `json:"id"`
	Title            string `json:"title"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
	Created, Updated time.Time
}

func (rl *Role) UseUnixTimestamp() {
	rl.CreatedAt = rl.Created.Unix()
	rl.UpdatedAt = rl.Updated.Unix()
}

func (rl *Role) FastScan() []interface{} {
	return []interface{}{
		&rl.ID,
		&rl.Title,
		&rl.Created,
		&rl.Updated,
	}
}