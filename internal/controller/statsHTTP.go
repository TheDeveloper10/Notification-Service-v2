package controller

import (
	"fmt"
	"notification-service/internal/dto"
	"notification-service/internal/service"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type StatsHTTP struct {
	clientSvc       *service.Client
	templateSvc     *service.Template
	notificationSvc *service.Notification

	executionTimesHTTP   map[string]*dto.ExecutionTimes
	executionTimesHTTPMu sync.Mutex

	startTime time.Time
}

func (ctrl *StatsHTTP) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"executionTimesHTTP": ctrl.executionTimesHTTP,
		"activeClientsHTTP":  ctrl.clientSvc.GetActiveClientsCount(),

		"cachedTemplatesCount": ctrl.templateSvc.GetCachedTemplatesCount(),
		"templatesCacheHits":   ctrl.templateSvc.GetTemplatesCacheHitStats(),

		"notifications": fiber.Map{
			"email": ctrl.notificationSvc.GetEmailStats(),
			"sms":   ctrl.notificationSvc.GetSMSStats(),
			"push":  ctrl.notificationSvc.GetPushStats(),
		},

		"upTime": time.Now().Unix() - ctrl.startTime.Unix(),
	})
}

func (ctrl *StatsHTTP) Middleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)

	key := fmt.Sprintf("%s %s", c.Method(), c.Route().Path)

	ctrl.executionTimesHTTPMu.Lock()
	defer ctrl.executionTimesHTTPMu.Unlock()

	executionTimes := ctrl.executionTimesHTTP[key]
	if executionTimes == nil {
		executionTimes = &dto.ExecutionTimes{}
	}

	executionTimes.TotalCalls++
	executionTimes.TotalTime += uint64(duration.Nanoseconds())

	ctrl.executionTimesHTTP[key] = executionTimes

	return err
}
