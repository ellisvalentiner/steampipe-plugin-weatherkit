package weatherkit

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableWeatherKitNextHourForecast() *plugin.Table {
	return &plugin.Table{
		Name:        "weatherkit_next_hour_forecast",
		Description: "WeatherKit Next Hour Forecast.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"latitude", "longitude"}),
			Hydrate:    listNextHourForecast,
		},
		Columns: []*plugin.Column{
			{
				Name:        "latitude",
				Type:        proto.ColumnType_STRING,
				Description: "A numeric value indicating the latitude of the coordinate between -90 and 90.",
				Transform:   transform.FromQual("latitude"),
			},
			{
				Name:        "longitude",
				Type:        proto.ColumnType_STRING,
				Description: "A numeric value indicating the longitude of the coordinate between -180 and 180.",
				Transform:   transform.FromQual("longitude"),
			},
			{
				Name:        "forecast_end",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time the forecast ends.",
			},
			{
				Name:        "forecast_start",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time the forecast starts.",
			},
			{
				Name:        "precipitation_chance",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The probability of precipitation during this minute.",
			},
			{
				Name:        "precipitation_intensity",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The precipitation intensity in millimeters per hour.",
			},
			{
				Name:        "start_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The start time of the minute.",
			},
			{
				Name:        "metadata",
				Type:        proto.ColumnType_JSON,
				Description: "Descriptive information about the weather data.",
			},
		},
	}
}

func listNextHourForecast(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	service, err := connect(ctx, d)
	if err != nil {
		logger.Error("Invalid credentials.")
		return nil, err
	}
	latitude := d.KeyColumnQuals["latitude"].GetStringValue()
	longitude := d.KeyColumnQuals["longitude"].GetStringValue()
	weather, _ := service.NextHourForecast(ctx, latitude, longitude)
	type Row struct {
		ForecastMinute
		ForecastEnd   string          `json:"forecastEnd,omitempty"`
		ForecastStart string          `json:"forecastStart,omitempty"`
		Metadata      WeatherMetadata `json:"metadata,omitempty"`
	}
	for _, minute := range weather.NextHourForecast.Minutes {
		row := Row{
			ForecastMinute: minute,
			ForecastEnd:    weather.NextHourForecast.ForecastEnd,
			ForecastStart:  weather.NextHourForecast.ForecastStart,
			Metadata:       weather.NextHourForecast.Metadata,
		}
		d.StreamListItem(ctx, row)
		if plugin.IsCancelled(ctx) {
			logger.Trace("CANCELLED!")
			return nil, nil
		}
	}
	return nil, nil
}
