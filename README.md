# Gym API

## Overview
This API is designed to check if a given user has been banned (blacklisted) on a specific server. The API is built using Golang's standard HTTP package and provides simple endpoints to verify user status.

## Project Structure
```
gym-api/
├── handler/
│   ├── check.go    # Handles user blacklist verification
│   ├── home.go     # Root endpoint handler
│   ├── gymapi.go   # Welcome endpoint handler
│   ├── ping.go     # Health check endpoint
│   └── server.go   # Server configuration and routing
├── main.go         # Application entry point
├── go.mod          # Go module definition
└── README.md       # This documentation
```

## API Responses
The `/gymapi/check` endpoint returns a JSON object with the following fields:

- `banned`: a boolean indicating whether the user is banned (`true` or `false`)
- `membership`: a string indicating the user's membership status (`not blacklisted` or `blacklisted`)
- `status`: an integer representing the HTTP status code (`200`)

### Example Response
```json
{
  "banned": false,
  "membership": "not blacklisted",
  "status": 200
}
```

## Endpoints
- `/`: Returns a "Hello World!" message (GET)
- `/gymapi`: Returns a "Welcome to the Gym API!" message (GET)
- `/gymapi/ping`: Health check endpoint, returns "GymAPI is running" (GET)
- `/gymapi/check?userid={user_id}`: Checks if a user is banned (POST)

## How to Run
1. Make sure you have Go installed on your system
2. Clone the repository
3. Run the application:
```bash
go run main.go
```
4. The server will start on port 9009 by default (or the port specified in the PORT environment variable)

## Implementation Details

### Server Configuration
The API uses Go's standard `net/http` package to handle HTTP requests:

```go
func StartServer() {
    // define routes
    http.HandleFunc("/", Home)
    http.HandleFunc("/gymapi", GymAPI)
    http.HandleFunc("/gymapi/ping", Ping)
    http.HandleFunc("/gymapi/check", Check)
    
    // Get port from environment or use default
    httpPort := os.Getenv("PORT")
    if httpPort == "" {
        httpPort = localPort // Default is 9009
    }
    
    // start server
    http.ListenAndServe(":"+httpPort, nil)
}
```

### User Validation
The API uses in-memory maps to validate users:

```go
// Map objects to mock database
var blacklistedUsers = map[string]bool{
    "101": true,
    "103": true,
    "105": true,
}

var notBlacklistedUsers = map[string]bool{
    "102": false,
    "104": false,
    "106": false,
}

// Returns true if userid is blacklisted
func validateUserid(userid string) bool {
    if blacklistedUsers[userid] {
        return true
    }
    if notBlacklistedUsers[userid] {
        return false
    }
    return false
}
```

## Testing the API
You can test the API using curl or any API testing tool:

```bash
# Test the root endpoint
curl http://localhost:9009/

# Test the welcome endpoint
curl http://localhost:9009/gymapi

# Test the ping endpoint
curl http://localhost:9009/gymapi/ping

# Test the check endpoint (for a blacklisted user)
curl -X POST "http://localhost:9009/gymapi/check?userid=101"

# Test the check endpoint (for a non-blacklisted user)
curl -X POST "http://localhost:9009/gymapi/check?userid=102"
```
