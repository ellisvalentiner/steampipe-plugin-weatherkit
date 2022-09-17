---
organization: ellisvalentiner
category: ["software development"]
icon_url: "/images/plugins/ellisvalentiner/weatherkit.svg"
brand_color: "#2684FF"
display_name: "WeatherKit"
short_name: "weatherkit"
description: "Steampipe plugin for querying weather from WeatherKit."
og_description: "Query weather with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/ellisvalentiner/weatherkit-social-graphic.png"
---

# WeatherKit + Steampipe

[WeatherKit](https://developer.apple.com/weatherkit/) is a service from Apple that provides state-of-the-art global weather data.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

## Examples

Get the current temperature and condition:

```sql
select
  as_of,
  temperature,
  condition_code
from
  weatherkit_current_weather
where
  latitude = 42.281
  and longitude = -83.743;
```

```
+---------------------------+--------------------+----------------+
| as_of                     | temperature        | condition_code |
+---------------------------+--------------------+----------------+
| 2022-07-12T15:19:15-04:00 | 27.350000381469727 | MostlyClear    |
+---------------------------+--------------------+----------------+
```

Get the precipitation forecast:

```sql
select
  forecast_start::date as forecast_date,
  precipitation_chance,
  precipitation_amount
from
  weatherkit_daily_forecast
where
  latitude=42.281
  and longitude=-83.743
order by
  forecast_date;
```

```
+---------------------+----------------------+----------------------+
| forecast_date       | precipitation_chance | precipitation_amount |
+---------------------+----------------------+----------------------+
| 2022-07-12 00:00:00 | 0.2800000011920929   | 0.9700000286102295   |
| 2022-07-13 00:00:00 | 0.5400000214576721   | 5.659999847412109    |
| 2022-07-14 00:00:00 | 0.009999999776482582 | <null>               |
| 2022-07-15 00:00:00 | 0.09000000357627869  | <null>               |
| 2022-07-16 00:00:00 | 0.4099999964237213   | 2.990000009536743    |
| 2022-07-17 00:00:00 | 0.6499999761581421   | 14                   |
| 2022-07-18 00:00:00 | 0.5799999833106995   | 8.430000305175781    |
| 2022-07-19 00:00:00 | 0.25999999046325684  | <null>               |
| 2022-07-20 00:00:00 | 0.3799999952316284   | 0.2800000011920929   |
| 2022-07-21 00:00:00 | 0.2800000011920929   | 0.03999999910593033  |
+---------------------+----------------------+----------------------+
```

## Documentation

- **[Table definitions & examples →](https://hub.steampipe.io/plugins/ellisvalentiner/weatherkit/tables)**

## Get started

### Install

Download and install the latest WeatherKit plugin:

```bash
steampipe plugin install ellisvalentiner/weatherkit
```

### Credentials

| Item        | Description                                                                                                                                                                                                 |
| :---------- |:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | Access to WeatherKit is included in the [Apple Developer Program](https://developer.apple.com/programs/). You’ll need to set up identifiers and keys before you can use the WeatherKit REST API (see [Request authentication for WeatherKit REST API](https://developer.apple.com/documentation/weatherkitrestapi/request_authentication_for_weatherkit_rest_api)). |

### Configuration

Installing the latest WeatherKit plugin will create a config file (`~/.steampipe/config/weatherkit.spc`) with a single connection named `weatherkit`:

```hcl
connection "weatherkit" {
    plugin    = "ellisvalentiner/weatherkit"

    # WeatherKit requires authorization using a signed developer token
    # You must either provide the information to generate a signed JSON web token (JWT) or supply a pre-generated JWT.
    # See the Apple Developer documentation
    # https://developer.apple.com/documentation/weatherkitrestapi/request_authentication_for_weatherkit_rest_api

    # Option 1: Generate an JWT
    # The 10-character key identifier from your developer account.
    # key_id = "STJY7HX969"

    # The service identifier.
    # service_id = "com.ellisvalentiner.weatherkit-client"

    # The Apple Developer Program (ADP) team identifier.
    # team_id = "JS4JVS2JBT"

    # Path to your private key for signing the JWT.
    # private_key_path = "~/.auth/AuthKey_STJY7HX969.p8"

    # Option 2: Use a pre-generated JWT
    # token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}

```

- `key_id` - The 10-character key identifier from your developer account.
- `service_id` - The service identifier.
- `team_id` - The Apple Developer Program (ADP) team identifier.
- `private_key_path` - Path to your private key for signing the JWT.
- `token` - Pre-generated JWT (optional).

#### Credentials from Environment Variables

The WeatherKit plugin will use the following environment variables, **only if other arguments (`key_id`, `service_id`, `team_id`, `private_key_path`, `token`) are not specified** in the connection:

```shell
export WEATHERKIT_KEY_ID="STJY7HX969"
export WEATHERKIT_SERVICE_ID="com.ellisvalentiner.weatherkit-client"
export WEATHERKIT_TEAM_ID="JS4JVS2JBT"
export WEATHERKIT_PRIVATE_KEY="~/.auth/AuthKey_STJY7HX969.p8"
export WEATHERKIT_TOKEN="eyJhbG..."
```

```hcl
connection "weatherkit" {
    plugin    = "ellisvalentiner/weatherkit"
}
```

## Get involved

- Open source: https://github.com/ellisvalentiner/steampipe-plugin-weatherkit
- Community: [Slack Channel](https://steampipe.io/community/join)
