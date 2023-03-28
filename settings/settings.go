package settings

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml
var settingsFile []byte

type DatbaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Settings struct {
	Port string        `yaml:"port"`
	DB   DatbaseConfig `yaml:"database"`
}

func New() (*Settings, error) {
	var s Settings
	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
