package scheduleHandlers

import (
	"sched/internal/core/ports"
	"sched/internal/dto"
	errmsg "sched/internal/errormsg"

	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	scheduleService ports.ScheduleService
}

// factory func
func NewHTTPHandler(scheduleService ports.ScheduleService) *HTTPHandler {
	return &HTTPHandler{
		scheduleService: scheduleService,
	}
}

/*
POST /v1/schedule, schedule a task
{
	when: '2022-02-02T03:03:00+07:00',
	http: {
		method: 'post',
		url: 'http://localhost:8443/health',
		headers: [{"key": "Content-Type", "value": "application/json"}],
		payload: "{\"name\":\"tj the boss\",\"person\":{\"age\": 21}}",
	}
}
*/

// HTTP handler to schedule a job
func (hdl *HTTPHandler) Schedule(c echo.Context) error {
	var req dto.ScheduleRequest
	if err := c.Bind(&req); err != nil {
		return errmsg.ErrorInvalidRequest(err.Error())
	}

	err := hdl.scheduleService.Schedule(req)

	return err
}
