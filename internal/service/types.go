package service

import (
	"notification-service/internal/dto"
	"notification-service/internal/util"
	"sync"
)

type notificationType struct {
	ContactInfo *string
	Body        *string
	SendFunc    func(*dto.Notification) util.StatusCode
}

type targetData struct {
	index           int
	notificationReq *dto.NotificationRequest
	target          *dto.NotificationTarget
	template        *dto.Template
}

type syncErrors struct {
	wg          sync.WaitGroup
	errorsChan  chan error
	errorsCount int
}

func (se *syncErrors) pushError(err error) {
	se.errorsChan <- err
	se.errorsCount++
}
