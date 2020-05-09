package main

import (
	"github.com/Gssssssssy/ns-stored/internal/pipeline"
	"github.com/Gssssssssy/ns-stored/internal/schedule"
	"github.com/Gssssssssy/ns-stored/pkg/log"
)

func main() {
	log.Infof(nil, "starting monitor ... ")
	// 启动 DataFilter (单例)
	pipeline.NewDataFilter()
	// 启动调度器
	scheduler := schedule.NewScheduler()
	defer scheduler.Close()
	scheduler.Start()
}
