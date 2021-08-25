package _interface

import (
	"context"
	"github.com/muhammadisa/nobita/constant"
)

type RWR interface {
	RedisLock(context.Context, int64, constant.ProcType) error
	RedisUnlock(context.Context, int64, constant.ProcType) error
	RedisCheck(context.Context, int64, constant.ProcType) bool
}
