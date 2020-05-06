package pipeline

import (
	"github.com/Gssssssssy/ns-stored/internal/task"
	"github.com/Gssssssssy/ns-stored/pkg/alarm/email"
	"github.com/pkg/errors"
	"sync"
)

var once sync.Once
var ResultPipeline *ResultFilter

func init() {
	ResultPipeline = NewResultFilter()
}

type ResultFilter struct {
	Chan chan *task.Result
	wg   sync.WaitGroup
}

func (tf *ResultFilter) Do() error {
	ret := <-tf.Chan
	if ret.IsAlarm {
		err := email.ServicePoint.Do(nil, ret)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (tf *ResultFilter) Add(job *task.Result) {
	tf.Chan <- job
}

func NewResultFilter() *ResultFilter {
	var tf *ResultFilter
	once.Do(func() {
		tf = &ResultFilter{}
	})
	return tf
}
