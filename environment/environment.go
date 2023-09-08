// MIT License
//
// Copyright (c) 2023 Alyx Holms
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
