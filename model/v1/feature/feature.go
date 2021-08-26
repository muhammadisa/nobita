package feature

import "time"

type Feature struct {
	ID               int64  `json:"id"`
	RoleID           int64  `json:"role_id"`
	Title            string `json:"title"`
	Name             string `json:"name"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
	Created, Updated time.Time
}

func (ft *Feature) UseUnixTimestamp() {
	ft.CreatedAt = ft.Created.Unix()
	ft.UpdatedAt = ft.Updated.Unix()
}

func (ft *Feature) FastScan() []interface{} {
	return []interface{}{
		&ft.ID,
		&ft.RoleID,
		&ft.Title,
		&ft.Name,
		&ft.Created,
		&ft.Updated,
	}
}
