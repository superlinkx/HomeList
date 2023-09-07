package environment

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	HostURL string `json:"host_url"`
}

func LoadConfig() (Config, error) {
	configFile := os.Getenv(ConfigFile.Env())

	if cfg, err := readConfig(configFile); err != nil {
		return cfg, fmt.Errorf("failed to read config file: %w", err)
	} else {
		envConfigFlag(&cfg.HostURL, HostUrl, cfg.HostURL, "Host URL")

		flag.Parse()
		return cfg, nil
	}
}

func envConfigFlag(p *string, name envVar, cfgVal string, usage string) {
	if val := os.Getenv(name.Env()); val != "" {
		cfgVal = val
	}

	flag.StringVar(p, name.Flag(), cfgVal, usage)
}

func readConfig(path string) (Config, error) {
	var config Config

	if path == "" {
		// If no path was specified, continue with empty configuration
		return config, nil
	} else if bytes, err := os.ReadFile(path); err != nil {
		return config, fmt.Errorf("could not read file at %s: %w", path, err)
	} else if err := json.Unmarshal(bytes, &config); err != nil {
		return config, fmt.Errorf("could not parse config at %s: %w", path, err)
	} else {
		return config, nil
	}
}
