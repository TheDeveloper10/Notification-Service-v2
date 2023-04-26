package config

const (
	serviceConfigPath = "config/service.yaml"
)

var (
	Master *MasterConfig
)

func InitMasterConfigs() {
	Master = &MasterConfig{}
	loadYAML(serviceConfigPath, &Master)
}
