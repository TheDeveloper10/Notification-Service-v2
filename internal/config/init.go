package config

const (
	serviceConfigPath = "config/service.yaml"
)

var (
	Master *MasterConfig
)

func InitConfigs() {
	Master = &MasterConfig{}
	loadYAML(serviceConfigPath, &Master)
}
