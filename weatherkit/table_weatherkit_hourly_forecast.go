package weatherkit

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"time"
)

func tableWeatherKitHourlyForecast() *plugin.Table {
	return &plugin.Table{
		Name:        "weatherkit_hourly_forecast",
		Description: "WeatherKit Hourly Forecast.",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.AllColumns([]string{"latitude", "longitude"}),
			KeyColumns: []*plugin.KeyColumn{
				{Name: "latitude"},
				{Name: "longitude"},
				{Name: "forecast_start", Operators: []string{"<", "<=", ">", ">="}},
			},
			Hydrate: listHourlyForecast,
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
				Name:        "cloud_cover",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The percentage of the sky covered with clouds during the period, from 0 to 1.",
			},
			{
				Name:        "condition_code",
				Type:        proto.ColumnType_STRING,
				Description: "An enumeration value indicating the condition at the time.",
			},
			{
				Name:        "daylight",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the hour starts during the day or night.",
			},
			{
				Name:        "forecast_start",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The starting date and time of the forecast.",
			},
			{
				Name:        "humidity",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The relative humidity at the start of the hour, from 0 to 1.",
			},
			{
				Name:        "precipitation_chance",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The chance of precipitation forecasted to occur during the hour, from 0 to 1.",
			},
			{
				Name:        "precipitation_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of precipitation forecasted to occur during the period.",
			},
			{
				Name:        "pressure",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The sea-level air pressure, in millibars.",
			},
			{
				Name:        "pressure_trend",
				Type:        proto.ColumnType_STRING,
				Description: "The direction of change of the sea-level air pressure.",
			},
			{
				Name:        "snowfall_intensity",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The rate at which snow crystals are falling, in millimeters per hour.",
			},
			{
				Name:        "temperature",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The temperature at the start of the hour, in degrees Celsius.",
			},
			{
				Name:        "temperature_apparent",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The feels-like temperature when considering wind and humidity, at the start of the hour, in degrees Celsius.",
			},
			{
				Name:        "temperature_dew_point",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The temperature at which relative humidity is 100% at the top of the hour, in degrees Celsius.",
			},
			{
				Name:        "uv_index",
				Type:        proto.ColumnType_INT,
				Description: "The level of ultraviolet radiation at the start of the hour.",
			},
			{
				Name:        "visibility",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The distance at which terrain is visible at the start of the hour, in meters.",
			},
			{
				Name:        "wind_direction",
				Type:        proto.ColumnType_INT,
				Description: "The direction of the wind at the start of the hour, in degrees.",
			},
			{
				Name:        "wind_gust",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The maximum wind gust speed during the hour, in kilometers per hour.",
			},
			{
				Name:        "wind_speed",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The wind speed at the start of the hour, in kilometers per hour.",
			},
			{
				Name:        "precipitation_amount",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The amount of precipitation forecasted to occur during period, in millimeters.",
			},
			{
				Name:        "metadata",
				Type:        proto.ColumnType_JSON,
				Description: "Descriptive information about the weather data.",
			},
		},
	}
}

func listHourlyForecast(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	service, err := connect(ctx, d)
	if err != nil {
		logger.Error("Invalid credentials.")
		return nil, err
	}
	latitude := d.KeyColumnQuals["latitude"].GetStringValue()
	longitude := d.KeyColumnQuals["longitude"].GetStringValue()
	params := make(map[string]string)
	hourlyStart := d.KeyColumnQuals["forecast_start"].GetTimestampValue()
	if hourlyStart != nil {
		params["hourlyStart"] = hourlyStart.AsTime().Format(time.RFC3339)
	}
	weather, _ := service.HourlyForecast(ctx, latitude, longitude, params)
	type Row struct {
		HourWeatherConditions
		Metadata WeatherMetadata `json:"metadata,omitempty"`
	}
	for _, hour := range weather.HourlyForecast.Hours {
		row := Row{
			HourWeatherConditions: hour,
			Metadata:              weather.HourlyForecast.Metadata,
		}
		d.StreamListItem(ctx, row)
		if plugin.IsCancelled(ctx) {
			logger.Trace("CANCELLED!")
			return nil, nil
		}
	}
	return nil, nil
}
