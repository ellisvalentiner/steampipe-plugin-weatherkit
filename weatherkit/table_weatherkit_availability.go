package weatherkit

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableWeatherKitAvailability() *plugin.Table {
	return &plugin.Table{
		Name:        "weatherkit_availability",
		Description: "WeatherKit Availability.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"latitude", "longitude"}),
			Hydrate:    listAvailability,
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
				Name:        "data_set",
				Type:        proto.ColumnType_STRING,
				Description: "The collection of weather information for a location.",
			},
		},
	}
}

func listAvailability(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	service, err := connect(ctx, d)
	if err != nil {
		logger.Error("Invalid credentials.")
		return nil, err
	}
	latitude := d.KeyColumnQuals["latitude"].GetStringValue()
	longitude := d.KeyColumnQuals["longitude"].GetStringValue()
	dataSet, _ := service.Availability(ctx, latitude, longitude)

	type Row struct {
		DataSet string `json:"dataSet,omitempty"`
	}

	for _, data := range dataSet {
		d.StreamListItem(ctx, Row{DataSet: data})
	}
	return nil, nil
}
