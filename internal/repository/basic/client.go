package basic

import (
	"notification-service/internal/client"
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type ClientRepository struct {
}

func (cr *ClientRepository) CreateClient(clientObj *dto.Client) util.StatusCode {
	_, err := client.Database.Exec(
		"insert into clients(id, secret, permissions) values(?, ?, ?)",
		clientObj.ID, clientObj.Secret, clientObj.Permissions.ToNumeric(),
	)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return util.StatusInternal
	}

	util.Logger.Info().Msgf("Created client %d", clientObj.ID)
	return util.StatusSuccess
}

func (cr *ClientRepository) GetClient(clientCredentials *dto.ClientCredentials) (*dto.ClientMetadata, util.StatusCode) {
	rows, err := client.Database.Query(
		"select permissions from clients where id=? and secret=?",
		clientCredentials.ID, clientCredentials.Secret,
	)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return nil, util.StatusInternal
	}
	defer rows.Close()

	if rows.Next() {
		var permissions util.PermissionsNumeric

		err := rows.Scan(&permissions)
		if err != nil {
			util.Logger.Error().Msg(err.Error())
			return nil, util.StatusInternal
		}

		util.Logger.Info().Msgf("Fetched client %d", clientCredentials.ID)
		return &dto.ClientMetadata{
			ID:          clientCredentials.ID,
			Permissions: permissions,
		}, util.StatusSuccess
	}

	util.Logger.Error().Msgf("Client %d not found", clientCredentials.ID)
	return nil, util.StatusNotFound
}
