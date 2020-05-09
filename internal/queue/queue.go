package queue

import (
	"github.com/Gssssssssy/ns-stored/internal/task"
)

var RQ *ResultQ

func init() {
	RQ = NewResultQ()
}

type ResultQ struct {
	ch chan *task.Result
}

func (rq *ResultQ) Submit(job *task.Result) {
	rq.ch <- job
}

func (rq *ResultQ) Get() *task.Result {
	job := <-rq.ch
	return job
}

func NewResultQ() *ResultQ {
	var rq *ResultQ
	rq = &ResultQ{
		ch: make(chan *task.Result),
	}
	return rq
}
