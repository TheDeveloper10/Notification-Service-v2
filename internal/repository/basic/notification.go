package basic

import (
	"notification-service/internal/client"
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type NotificationRepository struct {
}

func (nr *NotificationRepository) SaveNotification(notification *dto.Notification) (uint64, util.StatusCode) {
	result, err := client.Database.Exec(
		"insert into notifications(appId, templateId, contactInfo, title, message) values(?, ?, ?, ?, ?)",
		notification.AppID, notification.TemplateID, notification.ContactInfo, notification.Title, notification.Message,
	)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return 0, util.StatusInternal
	}

	id, err := result.LastInsertId()
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return 0, util.StatusInternal
	}

	util.Logger.Error().Msgf("Saved notification %d", id)
	return uint64(id), util.StatusSuccess
}
