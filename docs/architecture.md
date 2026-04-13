# SEC Baseball Hub - Architecture Documentation

## Technology Stack

**Backend:**
- Language: Go (Golang)
- Framework: net/http or Gin web framework
- Database: MySQL
- API Style: RESTful

**Frontend:**
- Framework: React
- Routing: React Router
- HTTP Client: Axios or Fetch API
- Styling: CSS Modules or Tailwind CSS

**Deployment:**
- Backend: Cloud platform (AWS, DigitalOcean, or Heroku)
- Frontend: Static hosting (Vercel, Netlify, or S3)
- Database: Managed MySQL instance

## System Architecture

```
┌─────────────────┐
│  React Frontend │
│   (Port 3000)   │
└────────┬────────┘
         │ HTTP/REST
         │
┌────────▼────────┐
│   Go Backend    │
│   (Port 8080)   │
└────────┬────────┘
         │ SQL
         │
┌────────▼────────┐
│  MySQL Database │
│   (Port 3306)   │
└─────────────────┘
```

## API Endpoints

### Teams Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/teams` | Get all SEC teams |
| GET | `/api/teams/:id` | Get specific team by ID |
| GET | `/api/teams/:id/roster` | Get team roster (all players) |
| GET | `/api/teams/:id/games` | Get all games for a team |
| POST | `/api/teams` | Create a new team (admin) |
| PUT | `/api/teams/:id` | Update team information (admin) |
| DELETE | `/api/teams/:id` | Delete a team (admin) |

### Players Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/players` | Get all players (with optional filters) |
| GET | `/api/players/:id` | Get specific player by ID |
| GET | `/api/players/top` | Get top players by various stats |
| POST | `/api/players` | Create a new player (admin) |
| PUT | `/api/players/:id` | Update player information (admin) |
| DELETE | `/api/players/:id` | Delete a player (admin) |

### Games Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/games` | Get all games (with optional filters) |
| GET | `/api/games/:id` | Get specific game by ID |
| GET | `/api/games/upcoming` | Get upcoming games only |
| GET | `/api/games/completed` | Get completed games only |
| POST | `/api/games` | Create a new game (admin) |
| PUT | `/api/games/:id` | Update game information (admin) |
| DELETE | `/api/games/:id` | Delete a game (admin) |

### Standings Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/standings` | Get current conference standings |
| GET | `/api/standings/:teamId` | Get standings info for specific team |
| PUT | `/api/standings/refresh` | Recalculate standings from game results |

### Statistics Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/stats/leaders` | Get statistical leaders across categories |
| GET | `/api/stats/teams/:id` | Get team statistics |
| GET | `/api/stats/players/:id` | Get player statistics |

## Detailed Endpoint Examples

### 1. GET `/api/teams` - Get All Teams

**Description:** Retrieves a list of all 16 SEC baseball teams with basic information.

**Request:**
```http
GET /api/teams HTTP/1.1
Host: api.secbaseballhub.com
Accept: application/json
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Alabama Crimson Tide",
      "abbreviation": "ALA",
      "logo_url": "https://cdn.secbaseballhub.com/logos/alabama.png",
      "conference_record": "12-6",
      "overall_record": "28-10",
      "conference_wins": 12,
      "conference_losses": 6,
      "overall_wins": 28,
      "overall_losses": 10,
      "win_percentage": 0.667,
      "home_stadium": "Sewell-Thomas Stadium",
      "location": "Tuscaloosa, AL"
    },
    {
      "id": 2,
      "name": "Arkansas Razorbacks",
      "abbreviation": "ARK",
      "logo_url": "https://cdn.secbaseballhub.com/logos/arkansas.png",
      "conference_record": "15-3",
      "overall_record": "32-6",
      "conference_wins": 15,
      "conference_losses": 3,
      "overall_wins": 32,
      "overall_losses": 6,
      "win_percentage": 0.833,
      "home_stadium": "Baum-Walker Stadium",
      "location": "Fayetteville, AR"
    }
  ],
  "count": 16
}
```

### 2. GET `/api/players?position=pitcher&limit=10&sort=era` - Get Top Pitchers

**Description:** Retrieves filtered list of players with query parameters for position, limit, and sorting.

**Request:**
```http
GET /api/players?position=pitcher&limit=10&sort=era&order=asc HTTP/1.1
Host: api.secbaseballhub.com
Accept: application/json
```

**Query Parameters:**
- `position` (optional): Filter by player position (pitcher, catcher, infielder, outfielder)
- `team_id` (optional): Filter by team ID
- `limit` (optional): Number of results to return (default: 50)
- `offset` (optional): Pagination offset (default: 0)
- `sort` (optional): Field to sort by (era, batting_avg, home_runs, etc.)
- `order` (optional): Sort order (asc, desc)

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 145,
      "first_name": "Hagen",
      "last_name": "Smith",
      "jersey_number": 33,
      "position": "Pitcher",
      "team_id": 2,
      "team_name": "Arkansas Razorbacks",
      "year": "Junior",
      "stats": {
        "games_played": 12,
        "innings_pitched": 78.1,
        "wins": 8,
        "losses": 1,
        "era": 2.18,
        "strikeouts": 95,
        "walks": 22,
        "whip": 1.02,
        "batting_average_against": 0.198
      }
    },
    {
      "id": 267,
      "first_name": "Hunter",
      "last_name": "Elliott",
      "jersey_number": 18,
      "position": "Pitcher",
      "team_id": 7,
      "team_name": "LSU Tigers",
      "year": "Sophomore",
      "stats": {
        "games_played": 10,
        "innings_pitched": 65.2,
        "wins": 6,
        "losses": 2,
        "era": 2.33,
        "strikeouts": 82,
        "walks": 18,
        "whip": 1.08,
        "batting_average_against": 0.205
      }
    }
  ],
  "count": 10,
  "total": 184
}
```

### 3. GET `/api/standings` - Get Conference Standings

**Description:** Retrieves current SEC conference standings with complete team records and rankings.

**Request:**
```http
GET /api/standings HTTP/1.1
Host: api.secbaseballhub.com
Accept: application/json
```

**Response:**
```json
{
  "success": true,
  "data": {
    "conference": "Southeastern Conference (SEC)",
    "season": "2025",
    "last_updated": "2025-04-13T10:30:00Z",
    "standings": [
      {
        "rank": 1,
        "team_id": 2,
        "team_name": "Arkansas Razorbacks",
        "abbreviation": "ARK",
        "logo_url": "https://cdn.secbaseballhub.com/logos/arkansas.png",
        "conference_wins": 15,
        "conference_losses": 3,
        "conference_win_pct": 0.833,
        "overall_wins": 32,
        "overall_losses": 6,
        "overall_win_pct": 0.842,
        "home_record": "18-2",
        "away_record": "10-3",
        "neutral_record": "4-1",
        "streak": "W5",
        "runs_scored": 387,
        "runs_allowed": 201
      },
      {
        "rank": 2,
        "team_id": 6,
        "team_name": "Florida Gators",
        "abbreviation": "FLA",
        "logo_url": "https://cdn.secbaseballhub.com/logos/florida.png",
        "conference_wins": 14,
        "conference_losses": 4,
        "conference_win_pct": 0.778,
        "overall_wins": 30,
        "overall_losses": 8,
        "overall_win_pct": 0.789,
        "home_record": "16-3",
        "away_record": "11-4",
        "neutral_record": "3-1",
        "streak": "W3",
        "runs_scored": 356,
        "runs_allowed": 215
      }
    ]
  }
}
```

## Error Responses

All endpoints follow a consistent error response format:

```json
{
  "success": false,
  "error": {
    "code": "NOT_FOUND",
    "message": "Team with ID 999 not found",
    "status": 404
  }
}
```

**Common Error Codes:**
- 400: Bad Request (invalid parameters)
- 401: Unauthorized (admin endpoints only)
- 404: Not Found
- 500: Internal Server Error

## Data Flow

1. **Client Request:** React frontend makes HTTP request to Go backend API
2. **Route Handling:** Go router matches request to handler function
3. **Business Logic:** Handler processes request, validates input
4. **Database Query:** Go application queries MySQL database
5. **Response Formation:** Data is formatted into JSON response
6. **Client Update:** React frontend receives data and updates UI

## Authentication (Future)

For admin endpoints (POST, PUT, DELETE), JWT-based authentication will be implemented:
- Login endpoint: `POST /api/auth/login`
- Protected routes require `Authorization: Bearer <token>` header
- Token validation middleware on admin routes

## CORS Configuration

Backend will be configured to accept requests from frontend origin:
- Development: `http://localhost:3000`
- Production: `https://secbaseballhub.com`

## Rate Limiting

To prevent abuse, API will implement rate limiting:
- 100 requests per minute per IP address
- 429 Too Many Requests response when exceeded
