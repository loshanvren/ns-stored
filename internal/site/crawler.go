package site

import (
	"context"
	"github.com/Gssssssssy/ns-onsale/internal/task"
)

const DefaultRetryTimes uint = 3

type Collector interface {
	Inquiry(ctx context.Context) (*task.Result, error)
}
