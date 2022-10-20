package weatherkit

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"net/http"
	"os"
	"strings"
)

const (
	baseUrl  = "weatherkit.apple.com"
	language = "en"
)

func connect(ctx context.Context, d *plugin.QueryData) (*Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "weatherkit"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*Client), nil
	}

	keyId := os.Getenv("WEATHERKIT_KEY_ID")
	serviceId := os.Getenv("WEATHERKIT_SERVICE_ID")
	teamId := os.Getenv("WEATHERKIT_TEAM_ID")
	privateKeyPath := os.Getenv("WEATHERKIT_PRIVATE_KEY")
	token := os.Getenv("WEATHERKIT_TOKEN")

	// Prefer config options given in Steampipe
	weatherKitConfig := GetConfig(d.Connection)
	var missingFields []string
	if weatherKitConfig.KeyId == nil {
		if keyId != "" {
			weatherKitConfig.KeyId = &keyId
		} else {
			missingFields = append(missingFields, "key_id")
		}
	}
	if weatherKitConfig.ServiceId == nil {
		if serviceId != "" {
			weatherKitConfig.ServiceId = &serviceId
		} else {
			missingFields = append(missingFields, "service_id")
		}
	}
	if weatherKitConfig.TeamId == nil {
		if teamId != "" {
			weatherKitConfig.TeamId = &teamId
		} else {
			missingFields = append(missingFields, "team_id")
		}
	}
	if weatherKitConfig.PrivateKeyPath == nil {
		if privateKeyPath != "" {
			weatherKitConfig.PrivateKeyPath = &privateKeyPath
		} else {
			missingFields = append(missingFields, "private_key_path")
		}
	}
	if weatherKitConfig.Token == nil && token != "" {
		weatherKitConfig.Token = &token
	}

	// If any fields are missing and a token is not supplied
	if len(missingFields) > 0 && weatherKitConfig.Token == nil {
		panic("\nInvalid configuration in ~/.steampipe/config/weatherkit.spc\nThe configuration is missing " +
			strings.Join(missingFields, ", ") +
			" and Token is undefined.\nEnsure key_id, service_id, team_id, and private_key_path are all defined or provide a pre-generated JWT")
	}

	// Make a new client that can hold the JWT
	client := NewClient(ctx, http.DefaultClient, &weatherKitConfig)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
