package main

import (
	"github.com/ellisvalentiner/steampipe-plugin-weatherkit/weatherkit"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: weatherkit.Plugin})
}
