package weatherkit

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableWeatherKitCurrentWeather() *plugin.Table {
	return &plugin.Table{
		Name:        "weatherkit_current_weather",
		Description: "WeatherKit Current Weather.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"latitude", "longitude"}),
			Hydrate:    getCurrentWeather,
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
				Name:        "as_of",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time.",
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
				Description: "A Boolean value indicating whether there is daylight.",
			},
			{
				Name:        "humidity",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The relative humidity, from 0 to 1.",
			},
			{
				Name:        "precipitation_intensity",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The precipitation intensity, in millimeters per hour.",
			},
			{
				Name:        "pressure",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The sea level air pressure, in millibars.",
			},
			{
				Name:        "pressure_trend",
				Type:        proto.ColumnType_STRING,
				Description: "The direction of change of the sea-level air pressure.",
			},
			{
				Name:        "temperature",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The current temperature, in degrees Celsius.",
			},
			{
				Name:        "temperature_apparent",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The feels-like temperature when factoring wind and humidity, in degrees Celsius.",
			},
			{
				Name:        "temperature_dew_point",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The temperature at which relative humidity is 100%, in Celsius.",
			},
			{
				Name:        "uv_index",
				Type:        proto.ColumnType_INT,
				Description: "The level of ultraviolet radiation.",
			},
			{
				Name:        "visibility",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The distance at which terrain is visible, in meters.",
			},
			{
				Name:        "wind_direction",
				Type:        proto.ColumnType_INT,
				Description: "The direction of the wind, in degrees.",
			},
			{
				Name:        "wind_gust",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The maximum wind gust speed, in kilometers per hour.",
			},
			{
				Name:        "wind_speed",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The wind speed, in kilometers per hour.",
			},
			{
				Name:        "metadata",
				Type:        proto.ColumnType_JSON,
				Description: "Descriptive information about the weather data.",
			},
		},
	}
}

func getCurrentWeather(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	service, err := connect(ctx, d)
	if err != nil {
		logger.Error("Invalid credentials.")
		return nil, err
	}
	latitude := d.KeyColumnQuals["latitude"].GetStringValue()
	longitude := d.KeyColumnQuals["longitude"].GetStringValue()
	weather, _ := service.CurrentWeather(ctx, latitude, longitude)
	type Row struct {
		CurrentWeatherData
		Metadata WeatherMetadata `json:"metadata,omitempty"`
	}
	row := Row{
		CurrentWeatherData: weather.CurrentWeather,
		Metadata:           weather.CurrentWeather.Metadata,
	}

	return row, nil
}
