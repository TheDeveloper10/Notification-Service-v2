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

	cache           map[uint64]*dto.CachedTemplate
	cacheMu         sync.RWMutex
	lastCleanupTime int64
	cacheHits       uint32
	cacheMisses     uint32
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
	if uint32(len(svc.cache)) >= config.Master.Service.Cache.TemplatesCacheLimit {
		util.Logger.Warn().Msg("Templates cache requires a cleanup")

		if time.Now().Unix()-svc.lastCleanupTime < int64(config.Master.Service.Cache.TemplatesCacheCleanupTime) {
			util.Logger.Warn().Msg("Templates cache clean timeout")
			return
		}

		svc.lastCleanupTime = time.Now().Unix()

		svc.cacheMu.Lock()
		defer svc.cacheMu.Unlock()

		toRemoveArr := []uint64{}
		for k, v := range svc.cache {
			if v.IsExpired() {
				toRemoveArr = append(toRemoveArr, k)
			}
		}

		for _, toRemove := range toRemoveArr {
			delete(svc.cache, toRemove)
		}

		util.Logger.Info().Msgf("Performed cache cleanup: removed %d caches", len(toRemoveArr))
	} else {
		svc.cacheMu.Lock()
		defer svc.cacheMu.Unlock()
	}

	svc.cache[template.ID] = &dto.CachedTemplate{
		Template:   template,
		ExpiryTime: time.Now().Unix() + int64(config.Master.Service.Cache.TemplatesCacheEntryExpiry),
	}
}

func (svc *Template) getTemplateFromCache(templateID uint64) *dto.Template {
	svc.cacheMu.RLock()
	if cachedTemplate, ok := svc.cache[templateID]; ok {
		svc.cacheHits++
		svc.cacheMu.RUnlock()

		if cachedTemplate.IsExpired() {
			svc.deleteTemplateFromCache(cachedTemplate.Template.ID)
			return nil
		}

		return cachedTemplate.Template
	} else {
		svc.cacheMisses++
	}
	svc.cacheMu.RUnlock()

	return nil
}

func (svc *Template) deleteTemplateFromCache(templateID uint64) {
	svc.cacheMu.Lock()
	delete(svc.cache, templateID)
	svc.cacheMu.Unlock()
}

func (svc *Template) GetCachedTemplatesCount() int {
	return len(svc.cache)
}

func (svc *Template) GetTemplatesCacheHits() uint32 {
	return svc.cacheHits
}

func (svc *Template) GetTemplatesCacheMisses() uint32 {
	return svc.cacheMisses
}
