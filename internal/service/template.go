package service

import (
	"notification-service/internal/config"
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"sync"
	"time"
)

type Template struct {
	templateRepo repository.ITemplate

	cache   map[uint64]*dto.CachedTemplate
	cacheMu sync.RWMutex
}

func (svc *Template) CreateTemplate(template *dto.Template) (uint64, util.StatusCode) {
	id, status := svc.templateRepo.CreateTemplate(template)
	if status != util.StatusSuccess {
		return id, status
	}

	template.ID = id

	svc.writeTemplateToCache(template)

	return id, status
}

func (svc *Template) UpdateTemplate(templateID uint64, template *dto.Template) util.StatusCode {
	status := svc.templateRepo.UpdateTemplate(templateID, template)
	if status != util.StatusSuccess {
		return status
	}

	template.ID = templateID

	svc.writeTemplateToCache(template)

	return status
}

func (svc *Template) GetTemplateByID(templateID uint64) (*dto.Template, util.StatusCode) {
	if template := svc.getTemplateFromCache(templateID); template != nil {
		return template, util.StatusSuccess
	}

	template, status := svc.templateRepo.GetTemplateByID(templateID)
	if status != util.StatusSuccess {
		return template, status
	}

	svc.writeTemplateToCache(template)

	return template, status
}

func (svc *Template) GetBulkTemplates(filter *dto.TemplateBulkFilter) ([]dto.Template, util.StatusCode) {
	return svc.templateRepo.GetBulkTemplates(filter)
}

func (svc *Template) DeleteTemplate(templateID uint64) util.StatusCode {
	svc.deleteTemplateFromCache(templateID)

	return svc.templateRepo.DeleteTemplate(templateID)
}

func (svc *Template) writeTemplateToCache(template *dto.Template) {
	svc.cacheMu.Lock()
	defer svc.cacheMu.Unlock()

	svc.cache[template.ID] = &dto.CachedTemplate{
		Template:   template,
		ExpiryTime: time.Now().Add(time.Second * time.Duration(config.Master.Service.Cache.TemplatesCacheEntryExpiry)).Unix(),
	}
}

func (svc *Template) getTemplateFromCache(templateID uint64) *dto.Template {
	svc.cacheMu.RLock()
	if cachedTemplate, ok := svc.cache[templateID]; ok {
		svc.cacheMu.RUnlock()

		if cachedTemplate.IsExpired() {
			svc.deleteTemplateFromCache(cachedTemplate.Template.ID)
			return nil
		}

		return cachedTemplate.Template
	}
	svc.cacheMu.RUnlock()

	return nil
}

func (svc *Template) deleteTemplateFromCache(templateID uint64) {
	svc.cacheMu.Lock()
	delete(svc.cache, templateID)
	svc.cacheMu.Unlock()
}
