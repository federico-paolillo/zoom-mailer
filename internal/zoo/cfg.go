package zoo

type ZooAppConfig struct {
	zooDirPath            string
	emailTemplateFilePath string
	sendlistFilePath      string
}

func NewZooAppConfig(
	zooDirPath string,
	emailTemplateFilePath string,
	sendlistFilePath string,
) *ZooAppConfig {

	return &ZooAppConfig{zooDirPath, emailTemplateFilePath, sendlistFilePath}

}
