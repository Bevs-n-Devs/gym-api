# Gym API for .NET and C#

## Overview
This API is designed to check if a given user has been banned on a specific server. The API is currently built using C# and is planned to be refactored into Golang.

## API Response
The API returns a JSON object with the following fields:

- banned: a boolean indicating whether the user is banned (`true` or `false`)
- membership: a string indicating the user's membership status (`not blacklisted` or `blacklisted`)
- status code: an integer representing the HTTP status code (`200` or `403`)

## Example Response
```
{
  "banned": false,
  "membership": "not blacklisted",
  "status code": 200
}
```