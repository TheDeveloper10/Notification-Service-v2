package mock

import (
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type NotificationRepository struct {
}

func (nr *NotificationRepository) SaveNotification(notification *dto.Notification) (uint64, util.StatusCode) {
	return 0, util.StatusSuccess
}

func (nr *NotificationRepository) GetBulkNotifications(filter *dto.NotificationBulkFilter) ([]dto.Notification, util.StatusCode) {
	return nil, util.StatusSuccess
}
