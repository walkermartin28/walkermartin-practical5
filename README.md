# SEC Baseball Hub

A comprehensive web application for tracking SEC baseball team standings, player statistics, and game schedules across all 16 Southeastern Conference teams.

## Overview

SEC Baseball Hub solves the problem of scattered SEC baseball data by providing fans, students, and alumni with a single, centralized platform to follow their favorite teams and players throughout the season.

**Key Features:**
- Live conference standings with team records
- Comprehensive player statistics and rankings
- Complete game schedules with filtering options
- Individual team pages with rosters and stats
- Mobile-responsive design for on-the-go access

## Tech Stack

### Backend
- **Language:** Go (Golang)
- **Framework:** net/http / Gin
- **Database:** MySQL 8.0+
- **API:** RESTful JSON API

### Frontend
- **Framework:** React 18+
- **Routing:** React Router v6
- **HTTP Client:** Axios
- **Styling:** CSS Modules / Tailwind CSS

### Database
- **Type:** MySQL
- **Version:** 8.0+
- **ORM:** database/sql or GORM

## Project Structure

```
walkermartin-practical4/
├── backend/          # Go backend application
├── frontend/         # React frontend application
├── docs/             # Project documentation
├── scripts/          # Database scripts and utilities
├── CLAUDE.md         # AI assistant context
└── README.md         # This file
```

## Getting Started

### Prerequisites

- **Go** 1.21 or higher ([Download](https://golang.org/dl/))
- **Node.js** 18+ and npm ([Download](https://nodejs.org/))
- **MySQL** 8.0+ ([Download](https://dev.mysql.com/downloads/))
- **Git** for version control

### Installation

#### 1. Clone the Repository

```bash
git clone <repository-url>
cd walkermartin-practical4
```

#### 2. Database Setup

Create the MySQL database and tables:

```bash
# Login to MySQL
mysql -u root -p

# Create database
CREATE DATABASE sec_baseball_hub;
USE sec_baseball_hub;

# Run schema creation
SOURCE scripts/seed-database.sql;
```

#### 3. Backend Setup

```bash
# Navigate to backend directory
cd backend

# Install dependencies
go mod download

# Create environment file
cp .env.example .env

# Edit .env with your database credentials
# DB_HOST=localhost
# DB_PORT=3306
# DB_USER=root
# DB_PASSWORD=your_password
# DB_NAME=sec_baseball_hub
# PORT=8080

# Run the backend server
go run cmd/server/main.go
```

The backend API will be running at `http://localhost:8080`

#### 4. Frontend Setup

```bash
# Navigate to frontend directory (from project root)
cd frontend

# Install dependencies
npm install

# Create environment file
cp .env.example .env

# Edit .env with API URL
# REACT_APP_API_URL=http://localhost:8080/api

# Start the development server
npm start
```

The frontend will be running at `http://localhost:3000`

### Verify Installation

1. Open browser to `http://localhost:3000`
2. You should see the SEC Baseball Hub homepage
3. Test API at `http://localhost:8080/api/teams` (should return JSON)

## Usage

### Viewing Standings

Navigate to the **Standings** page to see current SEC conference rankings, including:
- Conference and overall records
- Win percentages
- Team rankings

### Exploring Teams

Browse all 16 SEC teams on the **Teams** page, or click any team to view:
- Complete roster with player details
- Team statistics
- Upcoming game schedule

### Checking Player Stats

Visit the **Players** page to:
- Search for specific players
- Filter by position (pitchers, catchers, infielders, outfielders)
- Sort by various statistics (batting average, ERA, home runs, etc.)
- View individual player profiles

### Viewing Schedule

Access the **Schedule** page to:
- See all upcoming and completed games
- Filter by team or date range
- Check game times, locations, and results

## API Documentation

### Base URL
```
http://localhost:8080/api
```

### Key Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/teams` | Get all SEC teams |
| GET | `/teams/:id` | Get specific team |
| GET | `/teams/:id/roster` | Get team roster |
| GET | `/players` | Get all players |
| GET | `/players/:id` | Get specific player |
| GET | `/games` | Get all games |
| GET | `/standings` | Get conference standings |

**Example Request:**
```bash
curl http://localhost:8080/api/teams
```

**Example Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Alabama Crimson Tide",
      "abbreviation": "ALA",
      "conference_wins": 12,
      "conference_losses": 6,
      "overall_wins": 28,
      "overall_losses": 10
    }
  ]
}
```

For complete API documentation, see [docs/architecture.md](docs/architecture.md)

## Database Schema

The application uses four main tables:

1. **teams** - 16 SEC baseball teams
2. **players** - All players across all teams
3. **games** - Game schedules and results
4. **standings** - Conference standings (calculated)

For detailed schema information, see [docs/database-schema.md](docs/database-schema.md)

## Development

### Running Tests

**Backend:**
```bash
cd backend
go test ./...
```

**Frontend:**
```bash
cd frontend
npm test
```

### Building for Production

**Backend:**
```bash
cd backend
go build -o bin/server cmd/server/main.go
./bin/server
```

**Frontend:**
```bash
cd frontend
npm run build
# Build output in build/ directory
```

### Code Formatting

**Go:**
```bash
go fmt ./...
```

**JavaScript:**
```bash
npm run format
```

## Deployment

### Backend Deployment

1. **Choose a hosting platform:**
   - AWS EC2
   - DigitalOcean Droplet
   - Heroku
   - Google Cloud Platform

2. **Setup environment variables** on the hosting platform

3. **Deploy the Go binary:**
   ```bash
   go build -o server cmd/server/main.go
   ./server
   ```

4. **Configure MySQL database** (managed instance recommended)

### Frontend Deployment

1. **Build the production bundle:**
   ```bash
   npm run build
   ```

2. **Deploy to static hosting:**
   - **Vercel:** `vercel deploy`
   - **Netlify:** `netlify deploy --prod`
   - **AWS S3 + CloudFront:** Upload build folder to S3 bucket

3. **Update environment variables** with production API URL

### Database Deployment

1. Use a **managed MySQL service:**
   - AWS RDS
   - DigitalOcean Managed Databases
   - Google Cloud SQL
   - PlanetScale

2. **Run migrations** to create tables

3. **Seed initial data** (teams, initial games, etc.)

## Environment Variables

### Backend (.env)
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=sec_baseball_hub
PORT=8080
FRONTEND_URL=http://localhost:3000
```

### Frontend (.env)
```env
REACT_APP_API_URL=http://localhost:8080/api
```

## Project Documentation

- **[Project Proposal](docs/project-proposal.md)** - Features, user flows, and project scope
- **[Architecture](docs/architecture.md)** - API endpoints and system design
- **[Database Schema](docs/database-schema.md)** - Table structures and relationships
- **[CLAUDE.md](CLAUDE.md)** - AI assistant context and conventions

## Contributing

This is an independent project by Walker Martin for Web Application Integration course.

### Development Guidelines

1. Follow existing code patterns and conventions
2. Write tests for new features
3. Update documentation when making changes
4. Use meaningful commit messages
5. Keep components small and focused

## Troubleshooting

### Backend won't start
- Check MySQL is running: `mysql -u root -p`
- Verify database exists: `SHOW DATABASES;`
- Check environment variables in `.env`
- Verify port 8080 is not in use

### Frontend won't start
- Clear npm cache: `npm cache clean --force`
- Delete node_modules and reinstall: `rm -rf node_modules && npm install`
- Check port 3000 is available
- Verify API URL in frontend `.env`

### Database connection errors
- Verify MySQL credentials
- Check database name matches `.env`
- Ensure MySQL is running on correct port
- Test connection: `mysql -u root -p sec_baseball_hub`

### API requests failing
- Verify backend is running on port 8080
- Check CORS configuration allows frontend origin
- Test API directly: `curl http://localhost:8080/api/teams`
- Check browser console for errors

## License

This project is for educational purposes as part of the Web Application Integration course.

## Contact

**Author:** Walker Martin
**Course:** Web Application Integration
**Project:** Independent Project - SEC Baseball Hub

## Acknowledgments

- SEC Baseball for inspiration
- Web Application Integration course materials
- React and Go communities

---

**Current Version:** 1.0.0-alpha
**Last Updated:** April 2026
**Status:** In Development
