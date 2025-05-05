using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;

public class GymApiController
{
    public static void MapGymApiEndpoint(WebApplication app)
    {
        app.MapGet("/gymapi", () => Results.Ok(new { message = "Welcome to the Gym API" }));
    }
}
