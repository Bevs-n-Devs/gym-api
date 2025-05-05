var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

// Configure endpoints
HomeController.MapHomeEndpoint(app);
PingController.MapPingEndpoint(app);
GymApiController.MapGymApiEndpoint(app);
CheckController.MapCheckEndpoint(app);

app.Run();
