package _interface

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/role"
)

type RW interface {
	ReadRoles(context.Context) ([]role.Role, error)
	WriteRole(context.Context, role.Role) error
	EditRole(context.Context, role.Role) error
	DeleteRole(context.Context, int64) error
}
