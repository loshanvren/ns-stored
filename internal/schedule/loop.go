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
			jobs   = []task.Task{task.BestBuy}
			jobErr error
		)
		for _, job := range jobs {
			ctx := context.Background()
			log.Infof(ctx, "starting collector to crawl %s ...", job.String())
			jobErr = sdl.doCollect(ctx, job)
			if jobErr != nil {
				log.Errorf(ctx, "run collect job error: %s", jobErr.Error())
				return
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

func (sdl *Scheduler) doCollect(ctx context.Context, t task.Task) error {
	var (
		clt    site.Collector
		err    error
		result *task.Result
	)
	clt = makeFactory(t)
	result, err = clt.Inquiry(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	sdl.resultQueue.Submit(result)
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
