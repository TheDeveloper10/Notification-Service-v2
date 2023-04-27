package repository

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository/basic"
	"notification-service/internal/repository/mock"
	"notification-service/internal/util"
)

type INotification interface {
	SaveNotification(notification *dto.Notification) util.StatusCode
	// GetBulk
}

func NewBasicNotificationRepository() INotification {
	return &basic.NotificationRepository{}
}

func NewMockNotificationRepository() INotification {
	return &mock.NotificationRepository{}
}
