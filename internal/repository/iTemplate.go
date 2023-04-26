package repository

import (
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type ITemplate interface {
	CreateTemplate(template *dto.Template) (uint64, util.StatusCode)
	UpdateTemplate(templateID uint64, template *dto.Template) util.StatusCode
	GetTemplateByID(templateID uint64) (*dto.Template, util.StatusCode)
	// GetBulkTemplates(tem)
	DeleteTemplate(templateID uint64) util.StatusCode
}
