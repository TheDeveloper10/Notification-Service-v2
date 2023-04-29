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

func (t *Template) CreateTemplate(template *dto.Template) (uint64, util.StatusCode) {
	id, status := t.templateRepo.CreateTemplate(template)
	if status != util.StatusSuccess {
		return id, status
	}

	template.ID = id

	t.cacheMu.Lock()
	defer t.cacheMu.Unlock()

	t.cache[id] = template

	return id, status
}

func (t *Template) UpdateTemplate(templateID uint64, template *dto.Template) util.StatusCode {
	status := t.templateRepo.UpdateTemplate(templateID, template)
	if status != util.StatusSuccess {
		return status
	}

	t.cacheMu.Lock()
	defer t.cacheMu.Unlock()

	t.cache[templateID] = template

	return status
}

func (t *Template) GetTemplateByID(templateID uint64) (*dto.Template, util.StatusCode) {
	t.cacheMu.RLock()
	if template, ok := t.cache[templateID]; ok {
		t.cacheMu.RUnlock()
		return template, util.StatusSuccess
	}
	t.cacheMu.RUnlock()

	template, status := t.templateRepo.GetTemplateByID(templateID)
	if status != util.StatusSuccess {
		return template, status
	}

	t.cacheMu.Lock()
	defer t.cacheMu.Unlock()

	t.cache[templateID] = template

	return template, status
}

func (t *Template) GetBulkTemplates(filter *dto.TemplateBulkFilter) ([]dto.Template, util.StatusCode) {
	return t.templateRepo.GetBulkTemplates(filter)
}

func (t *Template) DeleteTemplate(templateID uint64) util.StatusCode {
	t.cacheMu.Lock()
	delete(t.cache, templateID)
	t.cacheMu.Unlock()

	return t.templateRepo.DeleteTemplate(templateID)
}
