package basic

import (
	"database/sql"
	"fmt"
	"notification-service/internal/client"
	"notification-service/internal/dto"
	"notification-service/internal/util"
	"strings"
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

func (nr *NotificationRepository) GetBulkNotifications(filter *dto.NotificationBulkFilter) ([]dto.Notification, util.StatusCode) {
	where := nr.whereConditionsFromFilter(filter)

	query := "select id, appId, templateId, contactInfo, title, message, sentTime from notifications"
	if len(where) > 0 {
		query = query + " where " + strings.Join(where, " and ")
	}
	query = query + fmt.Sprintf(" limit %d", filter.PerPage)

	rows, err := client.Database.Query(query)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return nil, util.StatusInternal
	}
	defer rows.Close()

	notifications := make([]dto.Notification, 0)
	for rows.Next() {
		notification, err := nr.scanNotification(rows)
		if err != nil {
			util.Logger.Error().Msg(err.Error())
			return nil, util.StatusInternal
		}

		notifications = append(notifications, *notification)
	}

	return notifications, util.StatusSuccess
}

func (nr *NotificationRepository) whereConditionsFromFilter(filter *dto.NotificationBulkFilter) []string {
	where := make([]string, 0)

	if filter.AppID != nil {
		where = append(where, fmt.Sprintf("appId=\"%s\"", *filter.AppID))
	}

	if filter.TemplateID != nil {
		where = append(where, fmt.Sprintf("templateId=%d", *filter.TemplateID))
	}

	if filter.StartTime != nil {
		where = append(where, fmt.Sprintf("sentTime>=%d", *filter.StartTime))
	}

	if filter.EndTime != nil {
		where = append(where, fmt.Sprintf("sentTime<=%d", *filter.EndTime))
	}

	if filter.LastNotificationID > 0 {
		where = append(where, fmt.Sprintf("id>%d", filter.LastNotificationID))
	}

	return where
}

func (nr *NotificationRepository) scanNotification(rows *sql.Rows) (*dto.Notification, error) {
	notification := dto.Notification{}

	err := rows.Scan(
		&notification.ID,
		&notification.AppID,
		&notification.TemplateID,
		&notification.ContactInfo,
		&notification.Title,
		&notification.Message,
		&notification.SentTime,
	)

	return &notification, err
}
