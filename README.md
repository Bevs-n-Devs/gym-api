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

## Endpoints
- `/`: returns a hello world message
- `/gymapi/ping`: checks if the API is up and running
- `/gymapi/`: returns a gym API message
- `/gymapi/check/{user_id}`: checks if a user is banned

## HTTP Requests Using ASP.NET Core and C#
```
var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();
```
This code creates a new instance of the WebApplication class using the `CreateBuilder` method, and then builds the app instance using the `Build` method.

### Creating Endpoints in the Controllers Directory
Once you have the app instance, you can start creating endpoints in the Controllers directory. Here's an example of how to create a simple GET endpoint: 
```
app.MapGet("/", () => "Hello World!");
```
This code maps a GET request to the root URL ("/") and returns the string "Hello World!".

### Different HTTP Request Methods
You can use these methods to declare routes that respond to different types of HTTP requests.
```
app.MapGet("/users", () => ...); // responds to GET /users
app.MapPost("/users", () => ...); // responds to POST /users
app.MapPut("/users/{id}", () => ...); // responds to PUT /users/{id}
app.MapDelete("/users/{id}", () => ...); // responds to DELETE /users/{id}
``` 