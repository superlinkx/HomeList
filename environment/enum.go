package environment

import "strings"

type envVar string

const (
	Prefix = "HL_"
)

const (
	ConfigFile envVar = "CONFIG_FILE"
	HostUrl    envVar = "HOST_URL"
)

func (s envVar) Env() string {
	switch s {
	case HostUrl:
		return Prefix + string(HostUrl)
	case ConfigFile:
		return Prefix + string(ConfigFile)
	default:
		return ""
	}
}

func (s envVar) Flag() string {
	switch s {
	case HostUrl:
		return strings.ToLower(string(HostUrl))
	case ConfigFile:
		return "flag_not_implemented"
	default:
		return ""
	}
}
