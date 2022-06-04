package bootstrap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sched/external"
	"sched/internal/core/ports"
	"sched/internal/dto"
	"time"
)

func TaskProcessor(scheduleRepository ports.ScheduleRepository) {
	fmt.Println("[INFO] Starting task prcessor")

	for {
		job, err := scheduleRepository.FetchTask()
		if err != nil {
			fmt.Println("Cannot fetch from Redis", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		if job == nil {
			time.Sleep(1 * time.Second)
			continue
		}

		if int64(job.Score) <= external.CurrentTime() {
			fmt.Println("Run ", *job.Member)
			scheduleRepository.DeleteTask(job)
			makeHttpRequest(job.Member)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func makeHttpRequest(member *string) {
	var scheduleRequest dto.ScheduleRequest

	err := json.Unmarshal([]byte(*member), &scheduleRequest)

	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest(scheduleRequest.Http.Method, scheduleRequest.Http.Url, bytes.NewBuffer([]byte(scheduleRequest.Http.Payload)))

	for _, header := range scheduleRequest.Http.Headers {
		req.Header.Set(header.Key, header.Value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
