# WeatherKit Plugin for Steampipe

> #### Note: WeatherKit REST API is in _beta_. Contributes are welcome to fix issues that may arise as the API changes or other improvements to the plugin.

## Quick start

Install the plugin with Steampipe:

```sh
steampipe plugin install ellisvalentiner/weatherkit
```

Get the current temperature and condition:

```sql
select as_of,
       temperature,
       condition_code
from weatherkit_current_weather
where latitude = '42.281'
  and longitude = '-83.743';
```

Get the precipitation forecast:

```sql
select forecast_start::date as forecast_date,
       precipitation_chance,
       precipitation_amount
from weatherkit_daily_forecast
where latitude = '42.281'
  and longitude = '-83.743'
order by forecast_date;
```

## Developing

Prerequisites:

* Steampipe
* Golang

Clone:

```sh
git clone https://github.com/ellisvalentiner/steampipe-plugin-weatherkit.git
cd steampipe-plugin-weatherkit
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
nano ~/.steampipe/config/weatherkit.spc
```

Try it!

```sh
steampipe query
> .inspect weatherkit
```

## Legal

Apple Weather and Weather are trademarks of Apple Inc.

[Data Sources](https://weatherkit.apple.com/legal-attribution.html)
