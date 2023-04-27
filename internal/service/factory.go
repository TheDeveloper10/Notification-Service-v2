package service

import "notification-service/internal/repository"

func NewTemplateService(templateRepo repository.ITemplate) *Template {
	return &Template{
		templateRepo: templateRepo,
	}
}
