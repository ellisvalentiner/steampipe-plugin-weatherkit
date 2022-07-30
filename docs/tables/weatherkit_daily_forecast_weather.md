# Table: weatherkit_daily_forecast

Get the daily forecast for the specified location.

The `weatherkit_daily_forecast` table can be used to query the daily forecast for the requested location.
**You must specify location** in the where or join clause using the `latitude` and `longitude` columns.

## Examples

### Get the daily forecast for Ann Arbor, MI

```sql
select
  *
from
  weatherkit_daily_forecast
where
  latitude='42.281'
  and longitude='-83.743';
```

### Get the daily temperature min/max forecast

```sql
select
  forecast_start,
  temperature_min,
  temperature_max
from
  weatherkit_daily_forecast
where
  latitude='42.281'
  and longitude='-83.743'
order by
  forecast_start;
```

### Get the daytime forecast

```sql
select
  r.humidity,
  r."windSpeed",
  r."cloudCover",
  r."forecastEnd",
  r."conditionCode",
  r."forecastStart",
  r."windDirection",
  r."precipitationType"
from
  weatherkit_daily_forecast,
  jsonb_to_record(daytime_forecast) r(
    humidity numeric,
    "windSpeed" numeric,
    "cloudCover" numeric,
    "forecastEnd" timestamp,
    "conditionCode" text,
    "forecastStart" timestamp,
    "windDirection" int,
    "precipitationType" text
  )
where
  latitude='42.281'
  and longitude='-83.743'
order by
  r."forecastStart";
```
