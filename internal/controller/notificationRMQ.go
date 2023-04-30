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
		return &dto.Error{Error: "Body must be in JSON"}, true
	} else if err := body.Validate(); err != nil {
		return &dto.Error{Error: err.Error()}, true
	}

	errs, status := ctrl.notificationSvc.SendNotifications(&body)
	if status == util.StatusSuccess {
		return nil, true
	} else if status == util.StatusNotFound || status == util.StatusInternal {
		return &dto.Error{Error: errs[0]}, true
	}

	return &dto.Errors{Errors: errs}, true
}
