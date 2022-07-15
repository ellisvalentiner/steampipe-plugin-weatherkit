package weatherkit

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func Plugin(_ context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-weatherkit",
		DefaultTransform: transform.FromGo(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"weatherkit_availability":       tableWeatherKitAvailability(),
			"weatherkit_current_weather":    tableWeatherKitCurrentWeather(),
			"weatherkit_daily_forecast":     tableWeatherKitDailyForecast(),
			"weatherkit_hourly_forecast":    tableWeatherKitHourlyForecast(),
			"weatherkit_next_hour_forecast": tableWeatherKitNextHourForecast(),
			"weatherkit_weather_alert":      tableWeatherKitWeatherAlert(),
		},
	}
	return p
}
