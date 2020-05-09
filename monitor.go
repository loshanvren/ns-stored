package main

import (
	"github.com/Gssssssssy/ns-stored/internal/pipeline"
	"github.com/Gssssssssy/ns-stored/internal/schedule"
	"github.com/Gssssssssy/ns-stored/pkg/log"
	"time"
)

func main() {
	log.Infof(nil, "starting monitor ... %s", time.Now().Format("2006-01-02 15:04:05"))
	// 启动 DataFilter (单例)
	pipeline.NewDataFilter()
	// 启动调度器
	scheduler := schedule.NewScheduler()
	defer scheduler.Close()
	scheduler.Start()
}
