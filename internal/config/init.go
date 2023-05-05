package config

import "notification-service/internal/util"

const (
	ServiceConfigPath    = "config/service.json"
	FirebaseAdminSDKPath = "config/firebase-adminsdk.json"
)

var (
	Master *MasterConfig
)

func InitConfigs() {
	Master = &MasterConfig{}
	loadConfig(ServiceConfigPath, &Master)

	util.Logger.Info().Msg("Loaded all configs")
}
