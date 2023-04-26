package mock

import (
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type TemplateRepository struct {
}

func (tr *TemplateRepository) CreateTemplate(template *dto.Template) (uint64, util.StatusCode) {
	return 2, util.StatusSuccess
}

func (tr *TemplateRepository) UpdateTemplate(templateID uint64, template *dto.Template) util.StatusCode {
	return util.StatusSuccess
}

func (tr *TemplateRepository) GetTemplateByID(templateID uint64) (*dto.Template, util.StatusCode) {
	return nil, util.StatusSuccess
}

func (tr *TemplateRepository) GetBulkTemplates(filter *dto.TemplateBulkFilter) ([]dto.Template, util.StatusCode) {
	return nil, util.StatusSuccess
}

func (tr *TemplateRepository) DeleteTemplate(templateID uint64) util.StatusCode {
	return util.StatusSuccess
}
