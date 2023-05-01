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
		return &dto.Error{Error: "Body must be in JSON"}, true
	} else if err := body.Validate(); err != nil {
		return &dto.Error{Error: err.Error()}, true
	}

	if body.ID == 0 {
		_, status := ctrl.templateSvc.CreateTemplate(&body)
		if status != util.StatusSuccess {
			return &dto.Error{Error: "Failed to create template"}, true
		}
	} else {
		status := ctrl.templateSvc.UpdateTemplate(body.ID, &body)
		if status != util.StatusSuccess {
			return &dto.Error{Error: "Failed to update template"}, true
		}
	}

	return nil, true
}

func (ctrl *TemplateRMQ) Get(request *amqp.Delivery) (any, bool) {
	body := dto.ID{}

	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		return &dto.Error{Error: "Body must be in JSON"}, true
	} else if err := body.ValidateTemplate(); err != nil {
		return &dto.Error{Error: err.Error()}, true
	}

	template, status := ctrl.templateSvc.GetTemplateByID(body.ID)
	if status == util.StatusSuccess {
		return template, true
	} else if status == util.StatusNotFound {
		return &dto.Error{Error: "Template not found"}, true
	} else {
		return &dto.Error{Error: "Failed to get template"}, true
	}
}

func (ctrl *TemplateRMQ) Delete(request *amqp.Delivery) (any, bool) {
	body := dto.ID{}

	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		return &dto.Error{Error: "Body must be in JSON"}, true
	} else if err := body.ValidateTemplate(); err != nil {
		return &dto.Error{Error: err.Error()}, true
	}

	status := ctrl.templateSvc.DeleteTemplate(body.ID)
	if status == util.StatusSuccess {
		return nil, true
	} else if status == util.StatusNotFound {
		return &dto.Error{Error: "Template not found"}, true
	} else {
		return &dto.Error{Error: "Failed to get template"}, true
	}
}

func (ctrl *TemplateRMQ) GetBulk(request *amqp.Delivery) (any, bool) {
	body := dto.TemplateBulkFilter{}

	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		return &dto.Error{Error: "Body must be in JSON"}, true
	} else if err := body.Validate(); err != nil {
		return &dto.Error{Error: err.Error()}, true
	}

	templates, status := ctrl.templateSvc.GetBulkTemplates(&body)
	if status == util.StatusSuccess {
		return templates, true
	} else if status == util.StatusNotFound {
		return &dto.Error{Error: "Template not found"}, true
	} else {
		return &dto.Error{Error: "Failed to get template"}, true
	}
}
