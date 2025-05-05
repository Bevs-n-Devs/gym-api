using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;

public class PingController
{
    public static void MapPingEndpoint(WebApplication app)
    {
        app.MapGet("/gymapi/ping", () => Results.Ok(new { message = "GymAPI is running" }));
    }
}