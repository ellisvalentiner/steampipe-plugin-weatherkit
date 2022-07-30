# Table: weatherkit_current_weather

Get the current weather conditions for the specified location.

The `weatherkit_current_weather` table can be used to query the current weather for the requested location.
**You must specify location** in the where or join clause using the `latitude` and `longitude` columns.

## Examples

### Get the current weather for Ann Arbor, MI

```sql
select
  *
from
  weatherkit_current_weather
where
  latitude='42.281'
  and longitude='-83.743';
```

### Get the temperature, humidity, and dew point

```sql
select
  temperature,
  humidity,
  temperature_dew_point
from
  weatherkit_current_weather
where
  latitude='42.281'
  and longitude='-83.743';
```

### Get the temperature in degrees Fahrenheit

```sql
select
  temperature*9/5 + 32.
from
  weatherkit_current_weather
where
  latitude='42.281'
  and longitude='-83.743';
```

### Get the wind direction, gust, and speed

```sql
select
  wind_direction,
  wind_gust,
  wind_speed
from
  weatherkit_current_weather
where
  latitude='42.281'
  and longitude='-83.743';
```
