package schedule

import (
	"context"
	"runtime"
	"sync"
)

type Scheduler struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	numCPUs uint32
	pools   *TaskHelper
}

func NewScheduler() *Scheduler {
	sdl := new(Scheduler)
	sdl.ctx, sdl.cancel = context.WithCancel(context.Background())
	sdl.numCPUs = 1
	if runtime.NumCPU() > 1 {
		sdl.numCPUs = uint32(runtime.NumCPU()) - 1
	}
	sdl.pools = new(TaskHelper)
	//sdl.cron =
	return sdl
}

func (sdl *Scheduler) Start() {
	// 生成定时任务
	// 后台启动采集程序
}

func (sdl *Scheduler) Close() {
	sdl.cancel()
	sdl.wg.Wait()
}

func (sdl *Scheduler) cronJob() {

}