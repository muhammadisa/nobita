package _interface

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/feature"
)

type RW interface{
	ReadFeature(context.Context, int64) (*feature.Feature, error)
}
