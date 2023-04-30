package config

import "notification-service/internal/util"

const (
	ServiceConfigPath    = "config/service.yaml"
	FirebaseAdminSDKPath = "config/firebase-adminsdk.json"
)

var (
	Master *MasterConfig
)

func InitConfigs() {
	Master = &MasterConfig{}
	loadYAML(ServiceConfigPath, &Master)

	util.Logger.Info().Msg("Loaded all configs")
}
