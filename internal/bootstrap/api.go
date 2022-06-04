package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"sched/config"
	scheduleHandlers "sched/internal/handlers/scheduleHandler"
	"sched/pkg/meta"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Api(scheduleHandler *scheduleHandlers.HTTPHandler) {
	e := newHTTPServer(
		scheduleHandler,
	)

	if err := e.Start(fmt.Sprintf("%s:%d", "0.0.0.0", config.Get().HTTPServer.Port)); err != nil {
		e.Logger.Info("shutting down the server")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer e.Shutdown(ctx)
}

func newHTTPServer(
	scheduleHandler *scheduleHandlers.HTTPHandler,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.POST("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "server is running...")
	})

	// TODO: add basic auth middleware

	v1 := e.Group("/v1")
	scheduleGroup := v1.Group("/schedule")
	scheduleGroup.POST("", scheduleHandler.Schedule)

	return e
}

func customHTTPErrorHandler(err error, c echo.Context) {
	m := &meta.MetaError{}

	if metaErr, ok := meta.IsError(err); ok {
		m = metaErr
	} else if he, ok := err.(*echo.HTTPError); ok {
		m = meta.NewError(he.Code).AppendMessage(-1000, strings.ToLower(he.Message.(string)))
	} else {
		m = meta.MetaErrorInternalServer.AppendError(-1000, err)
	}

	c.JSON(m.HttpStatus, map[string]interface{}{
		"code": m.Code,
		"msg":  m.Message,
	})
	return
}
