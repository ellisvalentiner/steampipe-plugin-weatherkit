package weatherkit

import (
	"context"
	"errors"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"net/http"
)

const (
	baseUrl  = "weatherkit.apple.com"
	language = "en"
)

func connect(_ context.Context, d *plugin.QueryData) (*Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "weatherkit"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*Client), nil
	}

	// Prefer config options given in Steampipe
	weatherKitConfig := GetConfig(d.Connection)
	if weatherKitConfig.KeyId == nil || weatherKitConfig.ServiceId == nil || weatherKitConfig.TeamId == nil || weatherKitConfig.PrivateKeyPath == nil {
		return nil, errors.New("invalid configuration from ~/.steampipe/config/weatherkit.spc")
	}

	// Make a new client that can hold the JWT
	client := NewClient(http.DefaultClient, weatherKitConfig)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
