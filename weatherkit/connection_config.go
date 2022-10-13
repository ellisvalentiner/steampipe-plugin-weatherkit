package weatherkit

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type weatherKitConfig struct {
	KeyId          *string `cty:"key_id"`
	ServiceId      *string `cty:"service_id"`
	TeamId         *string `cty:"team_id"`
	PrivateKeyPath *string `cty:"private_key_path"`
	Token          *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"key_id": {
		Type: schema.TypeString,
	},
	"service_id": {
		Type: schema.TypeString,
	},
	"team_id": {
		Type: schema.TypeString,
	},
	"private_key_path": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &weatherKitConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) weatherKitConfig {
	if connection == nil || connection.Config == nil {
		return weatherKitConfig{}
	}
	config, _ := connection.Config.(weatherKitConfig)
	return config
}
