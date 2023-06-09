package service

import (
	"notification-service/internal/dto"
	"notification-service/internal/util"
	"sync"
)

type notificationType struct {
	ContactInfo *string
	Message     *string
	SendFunc    func(*dto.Notification) util.StatusCode
	Stats       *dto.SuccessFailureStats
}

type targetData struct {
	index           int
	notificationReq *dto.NotificationRequest
	target          *dto.NotificationTarget
	template        *dto.Template
}

type syncErrors struct {
	wg       sync.WaitGroup
	errorsMu sync.Mutex
	errors   []string
}

func (se *syncErrors) addError(err string) {
	se.errorsMu.Lock()
	se.errors = append(se.errors, err)
	se.errorsMu.Unlock()
}

func newSyncErrors() *syncErrors {
	return &syncErrors{
		wg:       sync.WaitGroup{},
		errorsMu: sync.Mutex{},
		errors:   []string{},
	}
}
