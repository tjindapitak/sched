package ports

import "sched/internal/dto"

type ScheduleService interface {
	Schedule(req dto.ScheduleRequest) error
}
