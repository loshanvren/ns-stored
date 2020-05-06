package main

import "github.com/Gssssssssy/ns-stored/internal/schedule"

func main() {
	// 启动 Pipeline

	// 启动调度器
	scheduler := schedule.NewScheduler()
	defer scheduler.Close()
	scheduler.Start()
}