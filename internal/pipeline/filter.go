package pipeline

import (
	"github.com/Gssssssssy/ns-stored/internal/schedule"
	"github.com/Gssssssssy/ns-stored/internal/task"
	"github.com/Gssssssssy/ns-stored/pkg/alarm/email"
	"sync"
	"time"
)

var once sync.Once
var ResultPipeline *resultFilterPipeline

func init() {
	ResultPipeline = NewResultFilterPipeline()
}

type resultFilterPipeline struct {
	Chan chan *task.Result
	wg   sync.WaitGroup
}

// Do 永久阻塞，处理采集结果
func (rfp *resultFilterPipeline) Do() {
	go func() {
		for {
			rfp.wg.Add(1)
			ret := <-rfp.Chan
			if ret.IsAlarm {
				err := email.ServicePoint.Do(nil, ret)
				if err != nil {
					panic(err)
				}
			}
			time.Sleep(500 * time.Millisecond)
			rfp.wg.Done()
		}
	}()
	rfp.wg.Add(1)
	rfp.wg.Wait()
}

func (rfp *resultFilterPipeline) Add(job *task.Result) {
	rfp.Chan <- job
}

func NewResultFilterPipeline() *resultFilterPipeline {
	var rfp *resultFilterPipeline
	once.Do(func() {
		rfp = &resultFilterPipeline{
			Chan: schedule.ResultQueue,
		}
	})
	return rfp
}
