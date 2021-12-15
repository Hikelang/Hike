package languagepack

import (
	"gear/pkg/gearconfig"
)

func LoadLanguagePack(code string) map[string]string {
	switch code {
	case "ru":
		return Russian
	case "en":
		return English
	default:
		return English
	}
}

func GetCurrentPack() (map[string]string, error) {
	config, err := gearconfig.GetConfig()
	if err != nil {
		return map[string]string{}, err
	}
	return LoadLanguagePack(config.Language), nil
}
