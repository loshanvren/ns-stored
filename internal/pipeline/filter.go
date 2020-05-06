package pipeline

import (
	"github.com/Gssssssssy/ns-onsale/internal/task"
	"github.com/Gssssssssy/ns-onsale/pkg/alarm/email"
	"github.com/pkg/errors"
	"sync"
)

var once sync.Once
var ResultSet *ResultFilter

func init() {
	ResultSet = NewResultFilter()
}

type ResultFilter struct {
	Chan chan *task.Result
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

func NewResultFilter() *ResultFilter {
	var tf *ResultFilter
	once.Do(func() {
		tf = &ResultFilter{}
	})
	return tf
}
