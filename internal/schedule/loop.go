package schedule

import (
	"context"
	"fmt"
	"github.com/Gssssssssy/ns-stored/internal/queue"
	"github.com/Gssssssssy/ns-stored/internal/site"
	bestbuyCom "github.com/Gssssssssy/ns-stored/internal/site/bestbuy.com"
	"github.com/Gssssssssy/ns-stored/internal/task"
	"github.com/Gssssssssy/ns-stored/pkg/config"
	"github.com/Gssssssssy/ns-stored/pkg/log"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
	"runtime"
	"sync"
)

type Scheduler struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	numCPUs     uint32
	resultQueue *queue.ResultQ
}

func NewScheduler() *Scheduler {
	sdl := new(Scheduler)
	sdl.ctx, sdl.cancel = context.WithCancel(context.Background())
	sdl.numCPUs = 1
	if runtime.NumCPU() > 1 {
		sdl.numCPUs = uint32(runtime.NumCPU()) - 1
	}
	sdl.resultQueue = queue.RQ

	return sdl
}

func (sdl *Scheduler) Start() {
	// 生成定时任务
	c := cron.New()
	iv := 5
	if config.GetInt("inquiry_interval") != 0 {
		iv = config.GetInt("inquiry_interval")
	}
	spec := fmt.Sprintf(`*/%d * * * * *`, iv)
	// 后台启动采集
	err := c.AddFunc(spec, func() {
		var (
			jobs = []task.Task{task.BestBuy}
		)
		for _, job := range jobs {
			jobFailed := sdl.doCollect(context.Background(), job)
			if jobFailed != nil {
				log.Errorf(nil, "failed to collect product info: %v", jobFailed)
			}
		}
	})
	if err != nil {
		return
	}
	c.Start()
	// 永久阻塞
	select {}
}

func (sdl *Scheduler) Close() {
	sdl.cancel()
	sdl.wg.Wait()
}

func (sdl *Scheduler) doCollect(ctx context.Context, job task.Task) error {
	var (
		clt    site.Collector
		err    error
		result *task.Result
	)
	log.Infof(ctx, "run collector! %s data...", job.String())
	clt = makeFactory(job)
	result, err = clt.Inquiry(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	sdl.resultQueue.Submit(result)
	log.Infof(ctx, "done collector! %s data...", job.String())
	return nil
}

func makeFactory(t task.Task) site.Collector {
	switch t {
	case task.BestBuy:
		return bestbuyCom.NewCollector()
	default:
		return bestbuyCom.NewCollector()
	}
}
