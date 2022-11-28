package config

type ACULaConfig struct {
	TempDir  string
	AVType   string
	LogFile  string
	LogLevel string
	CertMode string
}

/* Загрузка конфигурации из json-файла
 */
func LoadConfigFromFile(filePath string) (c ACULaConfig, err error) {

	return c, nil
}

/* Создать шаблон конфигурационного файла
 */
func CreateDefaultConfig(filePath string) (err error) {

	return nil
}
