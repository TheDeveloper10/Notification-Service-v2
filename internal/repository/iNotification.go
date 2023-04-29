package repository

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository/basic"
	"notification-service/internal/repository/mock"
	"notification-service/internal/util"
)

type INotification interface {
	SaveNotification(notification *dto.Notification) (uint64, util.StatusCode)
	GetBulkNotifications(filter *dto.NotificationBulkFilter) ([]dto.Notification, util.StatusCode)
}

func NewBasicNotificationRepository() INotification {
	return &basic.NotificationRepository{}
}

func NewMockNotificationRepository() INotification {
	return &mock.NotificationRepository{}
}
