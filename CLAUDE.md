# CLAUDE.md - AI Assistant Context

This file provides context for AI coding assistants (like Claude, GPT, etc.) working on the SEC Baseball Hub project.

## Project Overview

**Project Name:** SEC Baseball Hub

**Description:** A web application that provides SEC baseball fans with a centralized platform to track team standings, player statistics, game schedules, and conference information across all 16 SEC baseball teams.

**Target Users:** SEC baseball fans, students, and alumni

**Key Value:** Consolidates scattered SEC baseball data from multiple sources into one clean, easy-to-use dashboard.

---

## Tech Stack

### Backend
- **Language:** Go (Golang)
- **Web Framework:** Standard library `net/http` or Gin framework
- **Database:** MySQL 8.0+
- **ORM/Query:** database/sql with mysql driver or GORM
- **API Style:** RESTful JSON API
- **Port:** 8080 (development)

### Frontend
- **Framework:** React 18+
- **Language:** JavaScript/JSX (TypeScript optional)
- **Build Tool:** Vite or Create React App
- **Routing:** React Router v6
- **HTTP Client:** Axios or Fetch API
- **Styling:** CSS Modules, Tailwind CSS, or styled-components
- **Port:** 3000 (development)

### Database
- **Type:** MySQL
- **Version:** 8.0 or higher
- **Port:** 3306
- **Character Set:** utf8mb4

### Development Tools
- **Version Control:** Git
- **Package Managers:** npm/yarn (frontend), go modules (backend)
- **Environment Variables:** .env files for configuration

---

## Project Structure

```
walkermartin-practical4/
├── backend/                    # Go backend application
│   ├── cmd/
│   │   └── server/
│   │       └── main.go        # Application entry point
│   ├── internal/
│   │   ├── handlers/          # HTTP request handlers
│   │   │   ├── teams.go
│   │   │   ├── players.go
│   │   │   ├── games.go
│   │   │   └── standings.go
│   │   ├── models/            # Data models/structs
│   │   │   ├── team.go
│   │   │   ├── player.go
│   │   │   ├── game.go
│   │   │   └── standing.go
│   │   ├── database/          # Database connection and queries
│   │   │   ├── db.go
│   │   │   └── migrations/
│   │   ├── middleware/        # HTTP middleware
│   │   │   ├── cors.go
│   │   │   └── logger.go
│   │   └── router/            # Route definitions
│   │       └── router.go
│   ├── configs/               # Configuration files
│   │   └── config.go
│   ├── go.mod                 # Go module definition
│   └── go.sum                 # Go dependencies
│
├── frontend/                  # React frontend application
│   ├── public/
│   │   └── index.html
│   ├── src/
│   │   ├── components/        # Reusable React components
│   │   │   ├── Header.jsx
│   │   │   ├── Footer.jsx
│   │   │   ├── TeamCard.jsx
│   │   │   ├── PlayerCard.jsx
│   │   │   ├── StandingsTable.jsx
│   │   │   └── GameSchedule.jsx
│   │   ├── pages/             # Page-level components
│   │   │   ├── Home.jsx
│   │   │   ├── Standings.jsx
│   │   │   ├── Teams.jsx
│   │   │   ├── TeamDetail.jsx
│   │   │   ├── Players.jsx
│   │   │   ├── PlayerDetail.jsx
│   │   │   └── Schedule.jsx
│   │   ├── services/          # API service layer
│   │   │   └── api.js
│   │   ├── utils/             # Helper functions
│   │   │   └── formatters.js
│   │   ├── App.jsx            # Root component
│   │   ├── index.jsx          # Entry point
│   │   └── App.css            # Global styles
│   ├── package.json
│   └── vite.config.js or package.json
│
├── docs/                      # Project documentation
│   ├── project-proposal.md
│   ├── architecture.md
│   └── database-schema.md
│
├── scripts/                   # Utility scripts
│   └── seed-database.sql      # Sample data for development
│
├── .env.example               # Example environment variables
├── .gitignore
├── CLAUDE.md                  # This file
└── README.md                  # Project README
```

---

## Key Conventions

### Backend (Go)

1. **Package Structure:**
   - Use `internal/` for private application code
   - Keep handlers thin - business logic in services
   - Models define struct types and validation

2. **Naming Conventions:**
   - Files: lowercase with underscores (e.g., `team_handler.go`)
   - Structs: PascalCase (e.g., `TeamHandler`)
   - Functions: PascalCase for exported, camelCase for private
   - Constants: PascalCase or SCREAMING_SNAKE_CASE

3. **Error Handling:**
   - Always check errors explicitly
   - Return errors up the call stack
   - Log errors with context
   - Return JSON error responses with consistent structure

4. **Database:**
   - Use prepared statements to prevent SQL injection
   - Handle connection pooling properly
   - Close database resources in defer statements
   - Use transactions for multi-step operations

5. **API Responses:**
   - Consistent JSON structure: `{"success": bool, "data": {}, "error": {}}`
   - Use proper HTTP status codes
   - Include timestamps in ISO 8601 format

### Frontend (React)

1. **Component Structure:**
   - Functional components with hooks
   - One component per file
   - Props validation with PropTypes or TypeScript

2. **Naming Conventions:**
   - Components: PascalCase (e.g., `TeamCard.jsx`)
   - Files: Match component name
   - CSS modules: `ComponentName.module.css`
   - Hooks: camelCase starting with "use" (e.g., `useTeams`)

3. **State Management:**
   - Use React hooks (useState, useEffect)
   - Context API for shared state if needed
   - Keep state as local as possible

4. **API Calls:**
   - Centralize API calls in `services/api.js`
   - Use async/await for promises
   - Handle loading and error states
   - Show user-friendly error messages

5. **Routing:**
   - Use React Router for navigation
   - Define routes in App.jsx
   - Use dynamic routes for detail pages (e.g., `/teams/:id`)

### Database

1. **Tables:**
   - Use lowercase with underscores: `teams`, `players`, `games`, `standings`
   - Primary keys named `id`
   - Foreign keys: `{table}_id` (e.g., `team_id`)
   - Include `created_at` and `updated_at` timestamps

2. **Queries:**
   - Use prepared statements
   - Index frequently queried columns
   - Optimize JOIN operations

---

## API Endpoints Quick Reference

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/teams` | Get all teams |
| GET | `/api/teams/:id` | Get team by ID |
| GET | `/api/teams/:id/roster` | Get team roster |
| GET | `/api/players` | Get all players (with filters) |
| GET | `/api/players/:id` | Get player by ID |
| GET | `/api/games` | Get all games (with filters) |
| GET | `/api/games/:id` | Get game by ID |
| GET | `/api/standings` | Get conference standings |

See `docs/architecture.md` for complete API documentation.

---

## Environment Variables

### Backend (.env)
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=sec_baseball_hub
PORT=8080
FRONTEND_URL=http://localhost:3000
```

### Frontend (.env)
```
REACT_APP_API_URL=http://localhost:8080/api
```

---

## Database Schema Quick Reference

**Tables:**
1. `teams` - 16 SEC baseball teams
2. `players` - All players across all teams
3. `games` - Schedule and results
4. `standings` - Conference standings (calculated)

**Relationships:**
- `teams` (1) → (many) `players` via `team_id`
- `teams` (1) → (many) `games` via `home_team_id` and `away_team_id`
- `teams` (1) → (1) `standings` via `team_id`

See `docs/database-schema.md` for complete schema documentation.

---

## Development Workflow

### Starting the Backend
```bash
cd backend
go mod download
go run cmd/server/main.go
```

### Starting the Frontend
```bash
cd frontend
npm install
npm start
```

### Database Setup
```bash
mysql -u root -p
CREATE DATABASE sec_baseball_hub;
USE sec_baseball_hub;
SOURCE scripts/seed-database.sql;
```

---

## Common Tasks for AI Assistants

### Creating New API Endpoints
1. Define route in `backend/internal/router/router.go`
2. Create handler function in appropriate handler file
3. Add database queries if needed
4. Update `docs/architecture.md` with new endpoint

### Adding New React Components
1. Create component file in `frontend/src/components/`
2. Import and use in appropriate page component
3. Add necessary props and state management
4. Style using chosen CSS methodology

### Database Changes
1. Update schema in `docs/database-schema.md`
2. Create migration SQL script
3. Update Go models in `backend/internal/models/`
4. Update affected handlers and queries

---

## Code Quality Standards

1. **Go:**
   - Run `go fmt` before committing
   - Use `golint` or `golangci-lint` for linting
   - Write tests for handlers and business logic
   - Document exported functions

2. **React:**
   - Use ESLint for code quality
   - Follow React best practices
   - Keep components small and focused
   - Write meaningful prop names

3. **General:**
   - Write descriptive commit messages
   - Comment complex logic
   - Keep functions small and single-purpose
   - Handle errors gracefully

---

## MVP Feature Checklist

- [ ] Live Conference Standings page
- [ ] Team Pages with rosters
- [ ] Game Schedule with filtering
- [ ] Player Statistics Dashboard
- [ ] Responsive design for mobile

---

## Testing

### Backend Testing
- Unit tests for handlers
- Integration tests for database operations
- Test API endpoints with sample data

### Frontend Testing
- Component tests with React Testing Library
- End-to-end tests with Cypress (optional)

---

## Deployment

### Backend
- Deploy to cloud platform (AWS EC2, DigitalOcean Droplet, Heroku)
- Use environment variables for configuration
- Set up MySQL database instance

### Frontend
- Build: `npm run build`
- Deploy to Vercel, Netlify, or S3 + CloudFront
- Configure API URL for production

---

## Important Notes for AI Assistants

1. **Always check existing code** before creating new files or functions
2. **Follow the established patterns** in the codebase
3. **Update documentation** when making significant changes
4. **Test endpoints** after creating or modifying them
5. **Handle errors** - never leave error cases unhandled
6. **Security first** - validate input, prevent SQL injection, sanitize data
7. **Keep it simple** - this is an MVP, avoid over-engineering
8. **Mobile-responsive** - all UI components should work on mobile devices

---

## Getting Help

- Review `docs/` folder for detailed documentation
- Check `README.md` for setup instructions
- Reference `docs/architecture.md` for API details
- See `docs/database-schema.md` for database structure

---

## Current Status

**Phase:** Initial setup and documentation
**Next Steps:** Backend setup, database creation, frontend scaffolding
