package weatherkit

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableWeatherKitDailyForecast() *plugin.Table {
	return &plugin.Table{
		Name:        "weatherkit_daily_forecast",
		Description: "WeatherKit Daily Forecast.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"latitude", "longitude"}),
			Hydrate:    listDailyForecast,
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
				Name:        "condition_code",
				Type:        proto.ColumnType_STRING,
				Description: "An enumeration value indicating the condition at the time.",
			},
			{
				Name:        "daytime_forecast",
				Type:        proto.ColumnType_JSON,
				Description: "The forecast between 7 AM and 7 PM for the day.",
			},
			{
				Name:        "forecast_end",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The ending date and time of the day.",
			},
			{
				Name:        "forecast_start",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The starting date and time of the day.",
			},
			{
				Name:        "max_uv_index",
				Type:        proto.ColumnType_INT,
				Description: "The maximum ultraviolet index value during the day.",
			},
			{
				Name:        "moon_phase",
				Type:        proto.ColumnType_STRING,
				Description: "The phase of the moon on the specified day.",
			},
			{
				Name:        "moonrise",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time of moonrise on the specified day.",
			},
			{
				Name:        "moonset",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time of moonset on the specified day.",
			},
			{
				Name:        "overnight_forecast",
				Type:        proto.ColumnType_JSON,
				Description: "The day part forecast between 7 PM and 7 AM for the overnight.",
			},
			{
				Name:        "precipitation_amount",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The amount of precipitation forecasted to occur during the day, in millimeters.",
			},
			{
				Name:        "precipitation_chance",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The chance of precipitation forecasted to occur during the day.",
			},
			{
				Name:        "precipitation_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of precipitation forecasted to occur during the day.",
			},
			{
				Name:        "snowfall_amount",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The depth of snow as ice crystals forecasted to occur during the day, in millimeters.",
			},
			{
				Name:        "solar_midnight",
				Type:        proto.ColumnType_STRING,
				Description: "The time when the sun is lowest in the sky.",
			},
			{
				Name:        "solar_noon",
				Type:        proto.ColumnType_STRING,
				Description: "The time when the sun is highest in the sky.",
			},
			{
				Name:        "sunrise",
				Type:        proto.ColumnType_STRING,
				Description: "The time when the top edge of the sun reaches the horizon in the morning.",
			},
			{
				Name:        "sunrise_astronomical",
				Type:        proto.ColumnType_STRING,
				Description: "The time when the sun is 18 degrees below the horizon in the morning.",
			},
			{
				Name:        "sunrise_civil",
				Type:        proto.ColumnType_STRING,
				Description: "The time when the sun is 6 degrees below the horizon in the morning.",
			},
			{
				Name:        "sunrise_nautical",
				Type:        proto.ColumnType_STRING,
				Description: "The time when the sun is 12 degrees below the horizon in the morning.",
			},
			{
				Name:        "sunset",
				Type:        proto.ColumnType_STRING,
				Description: "The time when the top edge of the sun reaches the horizon in the evening.",
			},
			{
				Name:        "sunset_astronomical",
				Type:        proto.ColumnType_STRING,
				Description: "The time when the sun is 18 degrees below the horizon in the evening.",
			},
			{
				Name:        "sunset_civil",
				Type:        proto.ColumnType_STRING,
				Description: "The time when the sun is 6 degrees below the horizon in the evening.",
			},
			{
				Name:        "sunset_nautical",
				Type:        proto.ColumnType_STRING,
				Description: "The time when the sun is 12 degrees below the horizon in the evening.",
			},
			{
				Name:        "temperature_max",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The maximum temperature forecasted to occur during the day, in degrees Celsius.",
			},
			{
				Name:        "temperature_min",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The minimum temperature forecasted to occur during the day, in degrees Celsius.",
			},
			{
				Name:        "metadata",
				Type:        proto.ColumnType_JSON,
				Description: "Descriptive information about the weather data.",
			},
		},
	}
}

func listDailyForecast(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	service, err := connect(ctx, d)
	if err != nil {
		logger.Error("Invalid credentials.")
		return nil, err
	}
	latitude := d.KeyColumnQuals["latitude"].GetStringValue()
	longitude := d.KeyColumnQuals["longitude"].GetStringValue()
	weather, _ := service.DailyForecast(ctx, latitude, longitude)
	type Row struct {
		DayWeatherConditions
		WeatherMetadata
	}
	for _, day := range weather.DailyForecast.Days {
		row := Row{
			DayWeatherConditions: day,
			WeatherMetadata:      weather.DailyForecast.Metadata,
		}
		d.StreamListItem(ctx, row)
		if plugin.IsCancelled(ctx) {
			logger.Trace("CANCELLED!")
			return nil, nil
		}
	}
	return nil, nil
}
