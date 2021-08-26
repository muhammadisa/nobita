package _interface

import (
	"context"
)

type RW interface {
	ReadFeature(context.Context, int64, string) error
}
