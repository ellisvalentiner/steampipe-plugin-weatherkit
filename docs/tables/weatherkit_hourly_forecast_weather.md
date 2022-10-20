# Table: weatherkit_hourly_forecast

Get the hourly forecast for the specified location.

The `weatherkit_hourly_forecast` table can be used to query the hourly forecast for the requested location.
**You must specify location** in the where or join clause using the `latitude` and `longitude` columns.

## Examples

### Get the hourly forecast for Ann Arbor, MI

```sql
select
  *
from
  weatherkit_hourly_forecast
where
  latitude=42.281
  and longitude=-83.743;
```

### Get the hourly precipitation forecast

```sql
select
  forecast_start,
  precipitation_amount,
  precipitation_chance,
  precipitation_type
from
  weatherkit_hourly_forecast
where
  latitude = 42.281
  and longitude = -83.743
order by
  forecast_start;
```

### Get the hourly wind direction, guest, and speed

```sql
select
  forecast_start,
  wind_direction,
  wind_gust,
  wind_speed
from
  weatherkit_hourly_forecast
where
  latitude = 42.281
  and longitude = -83.743
order by
  forecast_start;
```

### Get the hourly pressure change

```sql
select
  forecast_start,
  pressure - lag(pressure) over (order by forecast_start) as pressure_change
from
  weatherkit_hourly_forecast
where
  latitude = 42.281
  and longitude = -83.743
order by
  forecast_start;
```
