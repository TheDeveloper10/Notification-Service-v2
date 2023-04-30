package service

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"sync"
)

// TODO: add templates cache?

type Template struct {
	templateRepo repository.ITemplate
	cache        map[uint64]*dto.Template
	cacheMu      sync.RWMutex
}

func (svc *Template) CreateTemplate(template *dto.Template) (uint64, util.StatusCode) {
	id, status := svc.templateRepo.CreateTemplate(template)
	if status != util.StatusSuccess {
		return id, status
	}

	template.ID = id

	svc.cacheMu.Lock()
	defer svc.cacheMu.Unlock()

	svc.cache[id] = template

	return id, status
}

func (svc *Template) UpdateTemplate(templateID uint64, template *dto.Template) util.StatusCode {
	status := svc.templateRepo.UpdateTemplate(templateID, template)
	if status != util.StatusSuccess {
		return status
	}

	svc.cacheMu.Lock()
	defer svc.cacheMu.Unlock()

	svc.cache[templateID] = template

	return status
}

func (svc *Template) GetTemplateByID(templateID uint64) (*dto.Template, util.StatusCode) {
	svc.cacheMu.RLock()
	if template, ok := svc.cache[templateID]; ok {
		svc.cacheMu.RUnlock()
		return template, util.StatusSuccess
	}
	svc.cacheMu.RUnlock()

	template, status := svc.templateRepo.GetTemplateByID(templateID)
	if status != util.StatusSuccess {
		return template, status
	}

	svc.cacheMu.Lock()
	defer svc.cacheMu.Unlock()

	svc.cache[templateID] = template

	return template, status
}

func (svc *Template) GetBulkTemplates(filter *dto.TemplateBulkFilter) ([]dto.Template, util.StatusCode) {
	return svc.templateRepo.GetBulkTemplates(filter)
}

func (svc *Template) DeleteTemplate(templateID uint64) util.StatusCode {
	svc.cacheMu.Lock()
	delete(svc.cache, templateID)
	svc.cacheMu.Unlock()

	return svc.templateRepo.DeleteTemplate(templateID)
}
