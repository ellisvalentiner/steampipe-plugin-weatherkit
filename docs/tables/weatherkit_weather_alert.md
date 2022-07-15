# Table: weatherkit_weather_alert

List the weather alerts for the requested location.

The `weatherkit_weather_alert` table can be used to query information about severe weather alerts for the specified location.
**You must specify location** in the where or join clause using the `latitude` and `longitude` columns.

## Examples

### List weather alerts for Ann Arbor, MI

```sql
select *
from weatherkit_weather_alert
where latitude='42.281' and longitude='-83.743';
```

### List weather alert descriptions and expiration times for Austin, TX

```sql
select description, expire_time
from weatherkit_weather_alert
where latitude = '30.267'
  and longitude = '-97.743';
```
