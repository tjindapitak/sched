package ports

import (
	"sched/internal/dto"
)

type ScheduleRepository interface {
	AddToRedisSortedList(timestamp float64, payload string) error
	FetchTask() (*dto.ScheduledTask, error)
	DeleteTask(task *dto.ScheduledTask)
}
