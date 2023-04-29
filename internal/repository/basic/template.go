package basic

import (
	"database/sql"
	"fmt"
	"notification-service/internal/client"
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type TemplateRepository struct {
}

func (tr *TemplateRepository) CreateTemplate(template *dto.Template) (uint64, util.StatusCode) {
	result, err := client.Database.Exec(
		"insert into templates(bodyEmail, bodySMS, bodyPush, language, type) values(?, ?, ?, ?, ?)",
		template.Body.Email, template.Body.SMS, template.Body.Push, template.Language, template.Type,
	)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return 0, util.StatusInternal
	}

	id, err := result.LastInsertId()
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return 0, util.StatusInternal
	}

	util.Logger.Error().Msgf("Created template %d", id)
	return uint64(id), util.StatusSuccess
}

func (tr *TemplateRepository) UpdateTemplate(templateID uint64, template *dto.Template) util.StatusCode {
	result, err := client.Database.Exec(
		"update templates set bodyEmail=?, bodySMS=?, bodyPush=?, language=?, type=? where id=?",
		template.Body.Email, template.Body.SMS, template.Body.Push, template.Language, template.Type, templateID,
	)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return util.StatusInternal
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return util.StatusInternal
	} else if affectedRows <= 0 {
		util.Logger.Error().Msgf("Template %d not found", templateID)
		return util.StatusNotFound
	}

	util.Logger.Error().Msgf("Updated template %d", templateID)
	return util.StatusSuccess
}

func (tr *TemplateRepository) GetTemplateByID(templateID uint64) (*dto.Template, util.StatusCode) {
	rows, err := client.Database.Query(
		"select id, bodyEmail, bodySMS, bodyPush, language, type from templates where id=?",
		templateID,
	)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return nil, util.StatusInternal
	}
	defer rows.Close()

	if rows.Next() {
		template, err := tr.scanTemplate(rows)
		if err != nil {
			util.Logger.Error().Msg(err.Error())
			return nil, util.StatusInternal
		}

		util.Logger.Info().Msgf("Found template %d", templateID)
		return template, util.StatusSuccess
	}

	util.Logger.Error().Msgf("Template %d not found", templateID)
	return nil, util.StatusNotFound
}

func (tr *TemplateRepository) GetBulkTemplates(filter *dto.TemplateBulkFilter) ([]dto.Template, util.StatusCode) {
	query := "select id, bodyEmail, bodySMS, bodyPush, language, type from templates"

	if filter.LastTemplateID > 0 {
		query = query + fmt.Sprintf(" where id > %d", filter.LastTemplateID)
	}

	query = query + fmt.Sprintf(" limit %d", filter.PerPage)

	rows, err := client.Database.Query(query)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return nil, util.StatusInternal
	}
	defer rows.Close()

	templates := make([]dto.Template, 0)
	for rows.Next() {
		template, err := tr.scanTemplate(rows)
		if err != nil {
			util.Logger.Error().Msg(err.Error())
			return nil, util.StatusInternal
		}

		templates = append(templates, *template)
	}

	util.Logger.Info().Msgf("Fetched %d templates", len(templates))
	return templates, util.StatusSuccess
}

func (tr *TemplateRepository) DeleteTemplate(templateID uint64) util.StatusCode {
	result, err := client.Database.Exec(
		"delete from templates where id=?",
		templateID,
	)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return util.StatusInternal
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return util.StatusInternal
	} else if affectedRows <= 0 {
		util.Logger.Error().Msgf("Template %d not found", templateID)
		return util.StatusNotFound
	}

	util.Logger.Info().Msgf("Deleted template %d", templateID)
	return util.StatusSuccess
}

func (tr *TemplateRepository) scanTemplate(rows *sql.Rows) (*dto.Template, error) {
	template := dto.Template{Body: dto.TemplateBody{}}

	err := rows.Scan(
		&template.ID,
		&template.Body.Email,
		&template.Body.SMS,
		&template.Body.Push,
		&template.Language,
		&template.Type,
	)

	return &template, err
}
