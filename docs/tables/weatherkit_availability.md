# Table: weatherkit_availability

Determine the data sets available for the specified location.

The `weatherkit_availability` table can be used to query information about the data sets that are available for the specified location.
**You must specify location** in the where or join clause using the `latitude` and `longitude` columns.

## Examples

### List available data sets for Ann Arbor, MI

```sql
select
  *
from
  weatherkit_availability
where
  latitude=42.281
  and longitude=-83.743;
```
