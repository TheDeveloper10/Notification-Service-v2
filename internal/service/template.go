package service

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
)

type Template struct {
	templateRepo repository.ITemplate
}

func (t *Template) CreateTemplate(template *dto.Template) (uint64, util.StatusCode) {
	return t.templateRepo.CreateTemplate(template)
}

func (t *Template) UpdateTemplate(templateID uint64, template *dto.Template) util.StatusCode {
	return t.templateRepo.UpdateTemplate(templateID, template)
}

func (t *Template) GetTemplateByID(templateID uint64) (*dto.Template, util.StatusCode) {
	return t.templateRepo.GetTemplateByID(templateID)
}

func (t *Template) GetBulkTemplates(filter *dto.TemplateBulkFilter) ([]dto.Template, util.StatusCode) {
	return t.templateRepo.GetBulkTemplates(filter)
}

func (t *Template) DeleteTemplate(templateID uint64) util.StatusCode {
	return t.templateRepo.DeleteTemplate(templateID)
}
