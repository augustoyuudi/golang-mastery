# Requirements

- The endpoint should be exposed as `/api/time`
- The output should be in JSON
- A user can also request to return the current time in another timezone. /api/time?tz=America/New_York,Asia/Kolkata

- If the `tz` parameter is not provided in the URL, it should return the time in UTC:
```
{
  "current_time": "2021-08-09 11:18:06 +0000 UTC"
}
```
- In case of an invalid TZ database, the API should return the error message “invalid timezone” along with the status code 404
```
{
  "status": 404,
  "message": "invalid timezone"
}
```

API Response:
```
{
  "Asia/Kolkata": "2021-08-09 01:23:42 +0530 IST",
  "America/New_York": "2021-08-09 01:23:42 -0400 EDT"
}
```

> A list of TZ database names are available at https://en.wikipedia.org/wiki/List_of_tz_database_time_zones