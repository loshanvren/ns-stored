package schedule

import "github.com/Gssssssssy/ns-stored/internal/task"

var taskQueue chan *task.Task
var resultQueue chan *task.Result

func init() {
	taskQueue = make(chan *task.Task)
	resultQueue = make(chan *task.Result)
}

type TaskHelper struct{}

func (th *TaskHelper) Pull() *task.Task    { return <-taskQueue }
func (th *TaskHelper) Push(job *task.Task) { taskQueue <- job }

type ResultHelper struct{}

func (rh *ResultHelper) Pull() *task.Result    { return <-resultQueue }
func (rh *ResultHelper) Push(job *task.Result) { resultQueue <- job }