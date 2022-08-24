# Table: weatherkit_next_hour_forecast

Get the next hour forecast for the specified location.

The `weatherkit_next_hour_forecast` table can be used to query the forecast for the next hour for the requested location.
**You must specify location** in the where or join clause using the `latitude` and `longitude` columns.

## Examples

### Get the next hour forecast for Ann Arbor, MI

```sql
select
  *
from
  weatherkit_next_hour_forecast
where
  latitude='42.281'
  and longitude='-83.743';
```

### Get the next hour's precipitation chance & intensity

```sql
select
  start_time,
  precipitation_chance,
  precipitation_intensity
from
  weatherkit_next_hour_forecast
where
  latitude='42.281'
  and longitude='-83.743';
```
