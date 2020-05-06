package channel

import "github.com/Gssssssssy/ns-onsale/internal/task"

type TaskQueue interface {
	Push(job *task.Task) error
	Pop() (job *task.Task, ok bool)
}

type ResultQueue interface {
	Pop() (job *task.Result, ok bool)
}
