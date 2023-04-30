package controller

import (
	"encoding/json"
	"notification-service/internal/dto"
	"notification-service/internal/service"
	"notification-service/internal/util"

	amqp "github.com/rabbitmq/amqp091-go"
)

type NotificationRMQ struct {
	notificationSvc *service.Notification
}

func (ctrl *NotificationRMQ) Send(request *amqp.Delivery) (any, bool) {
	body := dto.NotificationRequest{}

	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		return "Body must be in JSON", true
	} else if err := body.Validate(); err != nil {
		return err.Error(), true
	}

	errs, status := ctrl.notificationSvc.SendNotifications(&body)
	if status == util.StatusSuccess {
		return nil, true
	} else if status == util.StatusNotFound || status == util.StatusInternal {
		return errs[0], true
	}

	return map[string]any{"errors": errs}, true
}
