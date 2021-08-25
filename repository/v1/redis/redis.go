package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/muhammadisa/nobita/constant"
	_interface "github.com/muhammadisa/nobita/repository/v1/redis/interface"
	"github.com/muhammadisa/nobita/util/dbc"
	"strconv"
)

type rwr struct {
	cmdable redis.Cmdable
}

func (r *rwr) RedisLock(ctx context.Context, accountID int64, lt constant.ProcType) error {
	var key = fmt.Sprintf(constant.RedisHaltKey, lt, accountID)
	return r.cmdable.Set(ctx, key, "true", 0).Err()
}

func (r *rwr) RedisUnlock(ctx context.Context, accountID int64, lt constant.ProcType) error {
	var key = fmt.Sprintf(constant.RedisHaltKey, lt, accountID)
	return r.cmdable.Set(ctx, key, "false", 0).Err()
}

func (r *rwr) RedisCheck(ctx context.Context, accountID int64, lt constant.ProcType) bool {
	var key = fmt.Sprintf(constant.RedisHaltKey, lt, accountID)
	res, err := r.cmdable.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	resultBool, err := strconv.ParseBool(res)
	if err != nil {
		return false
	}
	return resultBool
}

func NewRedisRepo(config dbc.Config) (_interface.RWR, error) {
	cmdable, err := dbc.OpenRedis(config)
	if err != nil {
		return nil, err
	}
	return &rwr{
		cmdable: cmdable,
	}, nil
}
