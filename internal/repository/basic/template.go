package basic

import (
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
		return 0, util.StatusError
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, util.StatusError
	}

	return uint64(id), util.StatusSuccess
}

func (tr *TemplateRepository) UpdateTemplate(templateID uint64, template *dto.Template) util.StatusCode {
	result, err := client.Database.Exec(
		"update templates set bodyEmail=?, bodySMS=?, bodyPush=?, language=?, type=? where id=?",
		template.Body.Email, template.Body.SMS, template.Body.Push, template.Language, template.Type, templateID,
	)
	if err != nil {
		return util.StatusError
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return util.StatusError
	} else if affectedRows <= 0 {
		return util.StatusNotFound
	}

	return util.StatusSuccess
}

func (tr *TemplateRepository) GetTemplateByID(templateID uint64) (*dto.Template, util.StatusCode) {
	rows, err := client.Database.Query(
		"select bodyEmail, bodySMS, bodyPush, language, type from templates where id=?",
		templateID,
	)
	if err != nil {
		return nil, util.StatusError
	}
	defer rows.Close()

	if rows.Next() {
		template := dto.Template{Body: dto.TemplateBody{}}

		err := rows.Scan(
			&template.Body.Email,
			&template.Body.SMS,
			&template.Body.Push,
			&template.Language,
			&template.Type,
		)
		if err != nil {
			return nil, util.StatusError
		}

		return &template, util.StatusSuccess
	}

	return nil, util.StatusNotFound
}

func (tr *TemplateRepository) GetBulkTemplates(filter *dto.TemplateBulkFilter) ([]dto.Template, util.StatusCode) {
	return nil, util.StatusSuccess
}

func (tr *TemplateRepository) DeleteTemplate(templateID uint64) util.StatusCode {
	result, err := client.Database.Exec(
		"delete from templates where id=?",
		templateID,
	)
	if err != nil {
		return util.StatusError
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return util.StatusError
	} else if affectedRows <= 0 {
		return util.StatusNotFound
	}

	return util.StatusSuccess
}
