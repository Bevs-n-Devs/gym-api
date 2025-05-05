using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using System.Collections.Generic;

public class CheckController
{
    // Define blacklisted and non-blacklisted user IDs
    private static readonly HashSet<string> BlacklistedUsers = new()
    {
        "101", "103", "105"
    };
    
    private static readonly HashSet<string> NotBlacklistedUsers = new()
    {
        "102", "104", "106"
    };
    
    public static void MapCheckEndpoint(WebApplication app)
    {
        app.MapPost("/gymapi/check", (HttpContext context) => 
        {
            string userid = context.Request.Query["userid"];
            // Check if the user is blacklisted
            bool isBanned = BlacklistedUsers.Contains(userid);
            
            if (isBanned)
            {
                return Results.Ok(new 
                { 
                    banned = true,
                    membership = "blacklisted",
                    status_code = 403
                });
            }
            else if (NotBlacklistedUsers.Contains(userid))
            {
                return Results.Ok(new 
                { 
                    banned = false,
                    membership = "not blacklisted",
                    status_code = 200
                });
            }
            else
            {
                // For any other user ID, return a default response
                return Results.Ok(new 
                { 
                    banned = false,
                    membership = "not blacklisted",
                    status_code = 200,
                    message = "User ID not in predefined lists, defaulting to not blacklisted"
                });
            }
        });
    }
}
