package config

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
}
