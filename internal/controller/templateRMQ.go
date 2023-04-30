package controller

import (
	"encoding/json"
	"notification-service/internal/dto"
	"notification-service/internal/service"
	"notification-service/internal/util"

	amqp "github.com/rabbitmq/amqp091-go"
)

type TemplateRMQ struct {
	templateSvc *service.Template
}

func (ctrl *TemplateRMQ) Write(request *amqp.Delivery) (any, bool) {
	body := dto.Template{}

	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		return "Body must be in JSON", true
	} else if err := body.Validate(); err != nil {
		return err.Error(), true
	}

	if body.ID == 0 {
		_, status := ctrl.templateSvc.CreateTemplate(&body)
		if status != util.StatusSuccess {
			return "Failed to create template", true
		}
	} else {
		status := ctrl.templateSvc.UpdateTemplate(body.ID, &body)
		if status != util.StatusSuccess {
			return "Failed to update template", true
		}
	}

	return nil, true
}

// TODO: delete, get, get bulk
