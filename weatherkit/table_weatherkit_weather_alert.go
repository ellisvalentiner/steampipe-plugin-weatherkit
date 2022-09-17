package weatherkit

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableWeatherKitWeatherAlert() *plugin.Table {
	return &plugin.Table{
		Name:        "weatherkit_weather_alert",
		Description: "WeatherKit Weather Alert.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"latitude", "longitude"}),
			Hydrate:    listWeatherAlert,
		},
		Columns: []*plugin.Column{
			{
				Name:        "latitude",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The latitude of the desired location.",
				Transform:   transform.FromQual("latitude"),
			},
			{
				Name:        "longitude",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The longitude of the desired location.",
				Transform:   transform.FromQual("longitude"),
			},
			{
				Name:        "area_id",
				Type:        proto.ColumnType_STRING,
				Description: "An official designation of the affected area.",
			},
			{
				Name:        "area_name",
				Type:        proto.ColumnType_STRING,
				Description: "A human-readable name of the affected area.",
			},
			{
				Name:        "start_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "A human-readable name of the affected area.",
			},
			{
				Name:        "certainty",
				Type:        proto.ColumnType_STRING,
				Description: "How likely the event is to occur.",
			},
			{
				Name:        "country_code",
				Type:        proto.ColumnType_STRING,
				Description: "The ISO code of the reporting country.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "A human-readable description of the event.",
			},
			{
				Name:        "details_url",
				Type:        proto.ColumnType_STRING,
				Description: "The URL to a page containing detailed information about the event.",
			},
			{
				Name:        "effective_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time the event went into effect.",
			},
			{
				Name:        "end_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the underlying weather event is projected to end.",
			},
			{
				Name:        "event_onset_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the underlying weather event is projected to start.",
			},
			{
				Name:        "expire_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the event expires.",
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "A unique identifier of the event.",
			},
			{
				Name:        "issued_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time that event was issued by the reporting agency.",
			},
			{
				Name:        "responses",
				Type:        proto.ColumnType_JSON,
				Description: "An array of recommended actions from the reporting agency.",
			},
			{
				Name:        "severity",
				Type:        proto.ColumnType_STRING,
				Description: "The level of danger to life and property.",
			},
			{
				Name:        "source",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the reporting agency.",
			},
			{
				Name:        "urgency",
				Type:        proto.ColumnType_STRING,
				Description: "An indication of urgency of action from the reporting agency.",
			},
			{
				Name:        "metadata",
				Type:        proto.ColumnType_JSON,
				Description: "Descriptive information about the weather data.",
			},
		},
	}
}

func listWeatherAlert(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	service, err := connect(ctx, d)
	if err != nil {
		logger.Error("Invalid credentials.")
		return nil, err
	}
	latitude := d.KeyColumnQuals["latitude"].GetDoubleValue()
	longitude := d.KeyColumnQuals["longitude"].GetDoubleValue()
	weather, _ := service.WeatherAlerts(ctx, latitude, longitude)
	logger.Debug("listWeatherAlert", "weather", weather)
	type Row struct {
		WeatherAlertSummary
		Metadata WeatherMetadata `json:"metadata,omitempty"`
	}
	for _, alert := range weather.WeatherAlerts.Alerts {
		row := Row{
			WeatherAlertSummary: alert,
			Metadata:            weather.WeatherAlerts.Metadata,
		}
		logger.Debug("listWeatherAlert", "row", row)
		d.StreamListItem(ctx, row)
		if plugin.IsCancelled(ctx) {
			logger.Trace("CANCELLED!")
			return nil, nil
		}
	}
	return nil, nil
}
