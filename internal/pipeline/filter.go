package pipeline

import (
	"github.com/Gssssssssy/ns-stored/internal/queue"
	"github.com/Gssssssssy/ns-stored/pkg/alarm/email"
	"github.com/Gssssssssy/ns-stored/pkg/log"
	"github.com/pkg/errors"
	"sync"
)

var once sync.Once
var DataFilter *dataFilter

type dataFilter struct {
	resultQ *queue.ResultQ
}

// Do 永久阻塞，处理采集结果
func (df *dataFilter) Do() (err error) {
	job := df.resultQ.Get()
	if job != nil {
		if job.IsAlarm {
			err = email.ServicePoint.Do(nil, job)
			if err != nil {
				return errors.WithStack(err)
			}
			log.Infof(nil, "send email succeed, job=%v", job)
		}
	}
	return nil
}

func NewDataFilter() *dataFilter {
	once.Do(func() {
		DataFilter = &dataFilter{resultQ: queue.RQ}
	})
	return DataFilter
}
