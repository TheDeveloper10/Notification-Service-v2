package controller

import (
	"notification-service/internal/dto"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type StatsHTTP struct {
	executionTimes   map[string]*dto.ExecutionTimes
	executionTimesMu sync.Mutex
}

func (ctrl *StatsHTTP) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"executionTimes": ctrl.executionTimes,
	})
}

func (ctrl *StatsHTTP) Middleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)

	path := c.Path()

	ctrl.executionTimesMu.Lock()
	defer ctrl.executionTimesMu.Unlock()

	executionTimes := ctrl.executionTimes[path]
	if executionTimes == nil {
		executionTimes = &dto.ExecutionTimes{}
	}
	executionTimes.TotalCalls++
	executionTimes.TotalTime += uint64(duration.Nanoseconds())

	ctrl.executionTimes[path] = executionTimes

	return err
}
