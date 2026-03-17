package internal

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// SDKConfig
type SDKConfig struct {
	CustomHeaders map[string]string `yaml:"custom_headers"`
	Logging       LoggingConfig     `yaml:"logging"`
}

// LoggingConfig
type LoggingConfig struct {
	ErrorLogFormat    string `yaml:"error_log_format"`
	EnableRequestLog  bool   `yaml:"enable_request_log"`
	EnableResponseLog bool   `yaml:"enable_response_log"`
}

var defaultConfig = SDKConfig{
	CustomHeaders: map[string]string{},
	Logging: LoggingConfig{
		ErrorLogFormat: "json",
	},
}

// LoadConfig searches for edgex-sdk.yaml in current directory and parent directories
func LoadConfig() (*SDKConfig, error) {
	configPath := findConfigFile("edgex-sdk.yaml")
	if configPath == "" {
		return &defaultConfig, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return &defaultConfig, nil
	}

	var config SDKConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return &defaultConfig, nil
	}

	if config.Logging.ErrorLogFormat == "" {
		config.Logging.ErrorLogFormat = "json"
	}

	return &config, nil
}

// findConfigFile searches for config file in current and parent directories
func findConfigFile(filename string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	// Try up to 5 parent directories
	for i := 0; i < 5; i++ {
		configPath := dir + "/" + filename
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}

		// Move to parent directory
		parent := dir + "/.."
		if absParent, err := filepath.Abs(parent); err == nil {
			if absParent == dir {
				// Reached root
				break
			}
			dir = absParent
		} else {
			break
		}
	}

	return ""
}
