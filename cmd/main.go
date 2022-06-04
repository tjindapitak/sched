package main

import (
	"os"
	"os/signal"
	"sched/config"
	"sched/external"
	bootstraps "sched/internal/bootstrap"
	"sched/internal/core/services"
	scheduleHandlers "sched/internal/handlers/scheduleHandler"
	"sched/internal/repositories"
	"syscall"
)

func main() {
	config.Init()

	// dependency registration
	redis := external.NewRedisClient()

	// repositories
	scheduleRepository := repositories.NewScheduleRespository(redis)

	// services
	scheduleService := services.NewScheduleService(scheduleRepository)

	// handlers
	scheduleHandler := scheduleHandlers.NewHTTPHandler(scheduleService)

	// go routine, API
	go bootstraps.Api(scheduleHandler)

	// go routing, tasks processor
	go bootstraps.TaskProcessor(scheduleRepository)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
