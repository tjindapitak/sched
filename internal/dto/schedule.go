package dto

import (
	"time"
)

type ScheduleHttpHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ScheduleHttpPayload struct {
	Method  string               `json:"method"`
	Url     string               `json:"url"`
	Headers []ScheduleHttpHeader `json:"headers"`
	Payload string               `json:"payload"`
}

type ScheduleRequest struct {
	ID   string              `json:"__id__"`
	When *time.Time          `json:"when"`
	Http ScheduleHttpPayload `json:"http"`
}

type ScheduledTask struct {
	Score  float64 `json:"Score"`
	Member *string `json:"Member"`
}
