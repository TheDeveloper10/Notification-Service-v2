package dto

import (
	"notification-service/internal/util"
)

type ClientMetadata struct {
	ID          string                  `json:"id"`
	Permissions util.PermissionsNumeric `json:"permissions"`
}
