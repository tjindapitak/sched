package repositories

import (
	"log"
	"sched/internal/core/ports"
	"sched/internal/dto"

	"github.com/go-redis/redis"
)

type scheduleRepository struct {
	redisClient *redis.Client
}

const taskRedisKey = "taskSet"

// factory func
func NewScheduleRespository(redisClient *redis.Client) ports.ScheduleRepository {
	return &scheduleRepository{
		redisClient: redisClient,
	}
}

// impl
func (repo *scheduleRepository) AddToRedisSortedList(timestamp float64, payload string) error {
	err := repo.redisClient.ZAdd(taskRedisKey, redis.Z{
		Score:  float64(timestamp),
		Member: payload,
	}).Err()

	if err != nil {
		log.Println("Error in adding job to redis: " + payload)
	}

	return err
}

func (repo *scheduleRepository) FetchTask() (*dto.ScheduledTask, error) {
	cmd := repo.redisClient.ZRangeWithScores(taskRedisKey, 0, 0)

	val, err := cmd.Result()
	if err != nil {
		log.Println("Error in getting job :", err.Error())
		return nil, err
	}

	// No tasks
	if len(val) == 0 {
		return nil, nil
	}

	member := val[0].Member.(string)

	return &dto.ScheduledTask{
		Score:  val[0].Score,
		Member: &member,
	}, nil
}

func (repo *scheduleRepository) DeleteTask(task *dto.ScheduledTask) {
	cmd := repo.redisClient.ZRem(taskRedisKey, *task.Member)

	if cmd.Err() != nil {
		log.Println("Cannot ZREM", cmd.Err().Error())
	}
}
