using Microsoft.AspNetCore.Builder;

public class HomeController
{
    public static void MapHomeEndpoint(WebApplication app)
    {
        app.MapGet("/", () => "Hello World!");
    }
}