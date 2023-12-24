package wordcounter

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ExcludePOSList []POS    `yaml:"exclude-pos-list"`
	KeepPOSList    []POS    `yaml:"keep-pos-list"`
	StopWords      []string `yaml:"stop-words"`
	Threshold      int      `yaml:"threshold"`
	UserDict       []string `yaml:"user-dict"`
}

func LoadConfig(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	configFile, err := io.ReadAll(f)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
