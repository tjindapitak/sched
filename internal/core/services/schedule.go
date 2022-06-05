package services

import (
	"encoding/json"
	"log"
	"sched/external"
	"sched/internal/core/ports"
	"sched/internal/dto"

	"github.com/google/uuid"
)

type scheduleSrv struct {
	scheduleRepository ports.ScheduleRepository
}

// factory func
func NewScheduleService(scheduleRepository ports.ScheduleRepository) ports.ScheduleService {
	return &scheduleSrv{
		scheduleRepository: scheduleRepository,
	}
}

// impl
func (service *scheduleSrv) Schedule(req dto.ScheduleRequest) error {
	uuid, _ := uuid.NewRandom()
	req.ID = uuid.String()
	b, err := json.Marshal(req)

	if err != nil {
		log.Println("Error converting json ", err.Error())
	}

	return service.scheduleRepository.AddToRedisSortedList(float64(external.GetEpocSecond(*req.When)), string(b))
}
