package auth

import "time"

type Verificator struct {
	Message          string `json:"message"`
	Accepted         bool   `json:"accepted"`
	ProfileCompleted bool   `json:"profile_completed"`
	AccountID        int64  `json:"account_id"`
	Token            string `json:"token"`
}

type Secret struct {
	Kind       string `json:"kind"`
	Identifier string `json:"identifier"`
	Code       string `json:"code"`
}

type Account struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	RoleID   int64  `json:"role_id"`
	DeviceID int64  `json:"device_id"`
	Timezone string `json:"timezone"`
}

type Contact struct {
	Email string
	Phone string
}

func (ct *Contact) FastScan() []interface{} {
	return []interface{}{
		&ct.Email,
		&ct.Phone,
	}
}

type Profile struct {
	ID               int64  `json:"id"`
	AccountID        int64  `json:"account_id"`
	LongLat          string `json:"long_lat"`
	FullName         string `json:"full_name"`
	BirthDay         string `json:"birth_day"`
	Gender           string `json:"gender"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
	Created, Updated time.Time
}

func (pr *Profile) UseUnixTimestamp() {
	pr.CreatedAt = pr.Created.Unix()
	pr.UpdatedAt = pr.Updated.Unix()
}

func (pr *Profile) FastScan() []interface{} {
	return []interface{}{
		&pr.ID,
		&pr.LongLat,
		&pr.AccountID,
		&pr.FullName,
		&pr.Gender,
		&pr.BirthDay,
		&pr.Created,
		&pr.Updated,
	}
}

type Status struct {
	Message string
	Sent    bool
}
