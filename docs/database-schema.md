# SEC Baseball Hub - Database Schema

## Database Overview

**Database Type:** MySQL 8.0+
**Character Set:** utf8mb4
**Collation:** utf8mb4_unicode_ci

This document describes the database schema for SEC Baseball Hub, including all tables, columns, data types, and relationships.

## Entity Relationship Diagram

```
┌─────────────┐         ┌─────────────┐
│    teams    │────┐    │   players   │
└─────────────┘    │    └─────────────┘
       │           │           │
       │           │           │
       │      ┌────▼───────────▼────┐
       │      │                     │
       └──────►       games         │
              │  (home_team_id,    │
              │   away_team_id)    │
              └─────────────────────┘
                      │
                      │
              ┌───────▼──────┐
              │  standings   │
              └──────────────┘
```

## Tables

### 1. `teams`

**Purpose:** Stores information about all 16 SEC baseball teams.

**Columns:**

| Column Name | Data Type | Constraints | Description |
|-------------|-----------|-------------|-------------|
| `id` | INT | PRIMARY KEY, AUTO_INCREMENT | Unique team identifier |
| `name` | VARCHAR(100) | NOT NULL, UNIQUE | Full team name (e.g., "Alabama Crimson Tide") |
| `abbreviation` | VARCHAR(10) | NOT NULL, UNIQUE | Team abbreviation (e.g., "ALA") |
| `logo_url` | VARCHAR(255) | NULL | URL to team logo image |
| `home_stadium` | VARCHAR(100) | NULL | Name of home stadium |
| `location` | VARCHAR(100) | NULL | City and state (e.g., "Tuscaloosa, AL") |
| `conference_wins` | INT | DEFAULT 0 | Number of conference wins |
| `conference_losses` | INT | DEFAULT 0 | Number of conference losses |
| `overall_wins` | INT | DEFAULT 0 | Total wins (all games) |
| `overall_losses` | INT | DEFAULT 0 | Total losses (all games) |
| `runs_scored` | INT | DEFAULT 0 | Total runs scored this season |
| `runs_allowed` | INT | DEFAULT 0 | Total runs allowed this season |
| `created_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |
| `updated_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | Last update timestamp |

**Indexes:**
- PRIMARY KEY on `id`
- UNIQUE INDEX on `name`
- UNIQUE INDEX on `abbreviation`

**Sample SQL:**
```sql
CREATE TABLE teams (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    abbreviation VARCHAR(10) NOT NULL UNIQUE,
    logo_url VARCHAR(255),
    home_stadium VARCHAR(100),
    location VARCHAR(100),
    conference_wins INT DEFAULT 0,
    conference_losses INT DEFAULT 0,
    overall_wins INT DEFAULT 0,
    overall_losses INT DEFAULT 0,
    runs_scored INT DEFAULT 0,
    runs_allowed INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

---

### 2. `players`

**Purpose:** Stores information about all SEC baseball players across all teams.

**Columns:**

| Column Name | Data Type | Constraints | Description |
|-------------|-----------|-------------|-------------|
| `id` | INT | PRIMARY KEY, AUTO_INCREMENT | Unique player identifier |
| `team_id` | INT | NOT NULL, FOREIGN KEY → teams(id) | Reference to player's team |
| `first_name` | VARCHAR(50) | NOT NULL | Player's first name |
| `last_name` | VARCHAR(50) | NOT NULL | Player's last name |
| `jersey_number` | INT | NULL | Player's jersey number |
| `position` | VARCHAR(30) | NOT NULL | Player position (Pitcher, Catcher, Infielder, Outfielder) |
| `year` | VARCHAR(20) | NULL | Academic year (Freshman, Sophomore, Junior, Senior) |
| `height` | VARCHAR(10) | NULL | Height (e.g., "6-2") |
| `weight` | INT | NULL | Weight in pounds |
| `hometown` | VARCHAR(100) | NULL | Player's hometown |
| `high_school` | VARCHAR(100) | NULL | High school or previous institution |
| `batting_avg` | DECIMAL(4,3) | NULL | Batting average (e.g., 0.315) |
| `home_runs` | INT | DEFAULT 0 | Number of home runs |
| `rbis` | INT | DEFAULT 0 | Runs batted in |
| `stolen_bases` | INT | DEFAULT 0 | Number of stolen bases |
| `era` | DECIMAL(4,2) | NULL | Earned run average (pitchers only) |
| `strikeouts` | INT | DEFAULT 0 | Strikeouts (pitchers: K's thrown, batters: K's taken) |
| `wins` | INT | DEFAULT 0 | Pitching wins |
| `saves` | INT | DEFAULT 0 | Pitching saves |
| `innings_pitched` | DECIMAL(5,1) | NULL | Innings pitched (pitchers only) |
| `games_played` | INT | DEFAULT 0 | Total games played |
| `created_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |
| `updated_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | Last update timestamp |

**Indexes:**
- PRIMARY KEY on `id`
- FOREIGN KEY on `team_id` REFERENCES `teams(id)` ON DELETE CASCADE
- INDEX on `team_id` for faster team roster queries
- INDEX on `position` for filtering by position
- INDEX on `batting_avg` for sorting
- INDEX on `era` for sorting

**Sample SQL:**
```sql
CREATE TABLE players (
    id INT AUTO_INCREMENT PRIMARY KEY,
    team_id INT NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    jersey_number INT,
    position VARCHAR(30) NOT NULL,
    year VARCHAR(20),
    height VARCHAR(10),
    weight INT,
    hometown VARCHAR(100),
    high_school VARCHAR(100),
    batting_avg DECIMAL(4,3),
    home_runs INT DEFAULT 0,
    rbis INT DEFAULT 0,
    stolen_bases INT DEFAULT 0,
    era DECIMAL(4,2),
    strikeouts INT DEFAULT 0,
    wins INT DEFAULT 0,
    saves INT DEFAULT 0,
    innings_pitched DECIMAL(5,1),
    games_played INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE
);

CREATE INDEX idx_team_id ON players(team_id);
CREATE INDEX idx_position ON players(position);
CREATE INDEX idx_batting_avg ON players(batting_avg);
CREATE INDEX idx_era ON players(era);
```

---

### 3. `games`

**Purpose:** Stores information about all SEC baseball games (scheduled and completed).

**Columns:**

| Column Name | Data Type | Constraints | Description |
|-------------|-----------|-------------|-------------|
| `id` | INT | PRIMARY KEY, AUTO_INCREMENT | Unique game identifier |
| `home_team_id` | INT | NOT NULL, FOREIGN KEY → teams(id) | Reference to home team |
| `away_team_id` | INT | NOT NULL, FOREIGN KEY → teams(id) | Reference to away team |
| `game_date` | DATE | NOT NULL | Date of the game |
| `game_time` | TIME | NULL | Scheduled start time |
| `location` | VARCHAR(100) | NULL | Stadium/venue name |
| `status` | ENUM('scheduled', 'in_progress', 'completed', 'postponed', 'cancelled') | DEFAULT 'scheduled' | Current game status |
| `home_score` | INT | NULL | Home team final score (null if not completed) |
| `away_score` | INT | NULL | Away team final score (null if not completed) |
| `innings` | INT | DEFAULT 9 | Number of innings played |
| `is_conference_game` | BOOLEAN | DEFAULT TRUE | Whether game counts toward conference standings |
| `attendance` | INT | NULL | Game attendance |
| `notes` | TEXT | NULL | Additional game notes |
| `created_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |
| `updated_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | Last update timestamp |

**Indexes:**
- PRIMARY KEY on `id`
- FOREIGN KEY on `home_team_id` REFERENCES `teams(id)` ON DELETE CASCADE
- FOREIGN KEY on `away_team_id` REFERENCES `teams(id)` ON DELETE CASCADE
- INDEX on `game_date` for date-based queries
- INDEX on `status` for filtering by game status
- INDEX on `home_team_id` for team schedule queries
- INDEX on `away_team_id` for team schedule queries

**Constraints:**
- CHECK constraint: `home_team_id != away_team_id` (team cannot play itself)

**Sample SQL:**
```sql
CREATE TABLE games (
    id INT AUTO_INCREMENT PRIMARY KEY,
    home_team_id INT NOT NULL,
    away_team_id INT NOT NULL,
    game_date DATE NOT NULL,
    game_time TIME,
    location VARCHAR(100),
    status ENUM('scheduled', 'in_progress', 'completed', 'postponed', 'cancelled') DEFAULT 'scheduled',
    home_score INT,
    away_score INT,
    innings INT DEFAULT 9,
    is_conference_game BOOLEAN DEFAULT TRUE,
    attendance INT,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (home_team_id) REFERENCES teams(id) ON DELETE CASCADE,
    FOREIGN KEY (away_team_id) REFERENCES teams(id) ON DELETE CASCADE,
    CHECK (home_team_id != away_team_id)
);

CREATE INDEX idx_game_date ON games(game_date);
CREATE INDEX idx_status ON games(status);
CREATE INDEX idx_home_team ON games(home_team_id);
CREATE INDEX idx_away_team ON games(away_team_id);
```

---

### 4. `standings`

**Purpose:** Stores calculated conference standings with additional metadata. This table can be periodically refreshed based on game results.

**Columns:**

| Column Name | Data Type | Constraints | Description |
|-------------|-----------|-------------|-------------|
| `id` | INT | PRIMARY KEY, AUTO_INCREMENT | Unique standings record identifier |
| `team_id` | INT | NOT NULL, FOREIGN KEY → teams(id), UNIQUE | Reference to team |
| `rank` | INT | NOT NULL | Current conference rank |
| `conference_win_pct` | DECIMAL(5,4) | NULL | Conference win percentage |
| `overall_win_pct` | DECIMAL(5,4) | NULL | Overall win percentage |
| `home_record` | VARCHAR(20) | NULL | Home record (e.g., "18-2") |
| `away_record` | VARCHAR(20) | NULL | Away record (e.g., "10-3") |
| `neutral_record` | VARCHAR(20) | NULL | Neutral site record |
| `streak` | VARCHAR(10) | NULL | Current streak (e.g., "W5", "L2") |
| `last_10` | VARCHAR(20) | NULL | Record in last 10 games (e.g., "8-2") |
| `games_back` | DECIMAL(3,1) | DEFAULT 0 | Games behind leader |
| `runs_per_game` | DECIMAL(4,2) | NULL | Average runs scored per game |
| `runs_allowed_per_game` | DECIMAL(4,2) | NULL | Average runs allowed per game |
| `updated_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | Last calculation timestamp |

**Indexes:**
- PRIMARY KEY on `id`
- FOREIGN KEY on `team_id` REFERENCES `teams(id)` ON DELETE CASCADE
- UNIQUE INDEX on `team_id` (one standings record per team)
- INDEX on `rank` for ordering

**Sample SQL:**
```sql
CREATE TABLE standings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    team_id INT NOT NULL UNIQUE,
    rank INT NOT NULL,
    conference_win_pct DECIMAL(5,4),
    overall_win_pct DECIMAL(5,4),
    home_record VARCHAR(20),
    away_record VARCHAR(20),
    neutral_record VARCHAR(20),
    streak VARCHAR(10),
    last_10 VARCHAR(20),
    games_back DECIMAL(3,1) DEFAULT 0,
    runs_per_game DECIMAL(4,2),
    runs_allowed_per_game DECIMAL(4,2),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE
);

CREATE INDEX idx_rank ON standings(rank);
```

---

## Relationships

### One-to-Many Relationships

1. **teams → players**
   - One team has many players
   - `players.team_id` references `teams.id`
   - Cascade delete: If a team is deleted, all its players are deleted

2. **teams → games (as home team)**
   - One team can be the home team in many games
   - `games.home_team_id` references `teams.id`
   - Cascade delete: If a team is deleted, all its games are deleted

3. **teams → games (as away team)**
   - One team can be the away team in many games
   - `games.away_team_id` references `teams.id`
   - Cascade delete: If a team is deleted, all its games are deleted

### One-to-One Relationships

1. **teams → standings**
   - One team has one standings record
   - `standings.team_id` references `teams.id`
   - Cascade delete: If a team is deleted, its standings record is deleted

---

## Sample Data

### Sample Teams
```sql
INSERT INTO teams (name, abbreviation, logo_url, home_stadium, location) VALUES
('Alabama Crimson Tide', 'ALA', 'https://cdn.example.com/alabama.png', 'Sewell-Thomas Stadium', 'Tuscaloosa, AL'),
('Arkansas Razorbacks', 'ARK', 'https://cdn.example.com/arkansas.png', 'Baum-Walker Stadium', 'Fayetteville, AR'),
('Auburn Tigers', 'AUB', 'https://cdn.example.com/auburn.png', 'Plainsman Park', 'Auburn, AL'),
('Florida Gators', 'FLA', 'https://cdn.example.com/florida.png', 'Florida Ballpark', 'Gainesville, FL'),
('Georgia Bulldogs', 'UGA', 'https://cdn.example.com/georgia.png', 'Foley Field', 'Athens, GA'),
('Kentucky Wildcats', 'UK', 'https://cdn.example.com/kentucky.png', 'Kentucky Proud Park', 'Lexington, KY'),
('LSU Tigers', 'LSU', 'https://cdn.example.com/lsu.png', 'Alex Box Stadium', 'Baton Rouge, LA'),
('Ole Miss Rebels', 'MISS', 'https://cdn.example.com/olemiss.png', 'Swayze Field', 'Oxford, MS'),
('Mississippi State Bulldogs', 'MSST', 'https://cdn.example.com/msstate.png', 'Dudy Noble Field', 'Starkville, MS'),
('Missouri Tigers', 'MIZ', 'https://cdn.example.com/missouri.png', 'Taylor Stadium', 'Columbia, MO'),
('Oklahoma Sooners', 'OU', 'https://cdn.example.com/oklahoma.png', 'L. Dale Mitchell Park', 'Norman, OK'),
('South Carolina Gamecocks', 'SC', 'https://cdn.example.com/southcarolina.png', 'Founders Park', 'Columbia, SC'),
('Tennessee Volunteers', 'TENN', 'https://cdn.example.com/tennessee.png', 'Lindsey Nelson Stadium', 'Knoxville, TN'),
('Texas A&M Aggies', 'TAMU', 'https://cdn.example.com/texasam.png', 'Blue Bell Park', 'College Station, TX'),
('Texas Longhorns', 'TEX', 'https://cdn.example.com/texas.png', 'UFCU Disch-Falk Field', 'Austin, TX'),
('Vanderbilt Commodores', 'VU', 'https://cdn.example.com/vanderbilt.png', 'Hawkins Field', 'Nashville, TN');
```

### Sample Players
```sql
INSERT INTO players (team_id, first_name, last_name, jersey_number, position, year, batting_avg, home_runs, rbis) VALUES
(2, 'Kendall', 'Diggs', 7, 'Outfielder', 'Junior', 0.342, 12, 45),
(2, 'Hagen', 'Smith', 33, 'Pitcher', 'Junior', NULL, 0, 0),
(6, 'Jac', 'Caglianone', 35, 'Pitcher', 'Senior', 0.328, 18, 52);
```

### Sample Games
```sql
INSERT INTO games (home_team_id, away_team_id, game_date, game_time, location, status, home_score, away_score, is_conference_game) VALUES
(2, 1, '2025-04-15', '18:00:00', 'Baum-Walker Stadium', 'scheduled', NULL, NULL, TRUE),
(6, 7, '2025-04-16', '19:00:00', 'Florida Ballpark', 'scheduled', NULL, NULL, TRUE),
(2, 6, '2025-04-10', '18:30:00', 'Baum-Walker Stadium', 'completed', 7, 3, TRUE);
```

---

## Database Maintenance

### Updating Standings
Standings should be recalculated after games are completed:
```sql
-- Update team records
UPDATE teams t
SET
    conference_wins = (SELECT COUNT(*) FROM games WHERE (home_team_id = t.id AND home_score > away_score AND is_conference_game = TRUE) OR (away_team_id = t.id AND away_score > home_score AND is_conference_game = TRUE)),
    conference_losses = (SELECT COUNT(*) FROM games WHERE (home_team_id = t.id AND home_score < away_score AND is_conference_game = TRUE) OR (away_team_id = t.id AND away_score < home_score AND is_conference_game = TRUE));
```

### Backup Strategy
- Daily automated backups
- Transaction logs for point-in-time recovery
- Retain backups for 30 days

---

## Future Enhancements

Potential schema additions for future versions:
- `game_stats` table for detailed box scores
- `player_game_stats` table for individual game performances
- `seasons` table for multi-year data
- `users` table for authentication and favorites
- `pitching_stats` and `batting_stats` separate tables for more detailed analytics
