-- SEC Baseball Hub Database Schema
-- Database: sec_baseball
-- Character Set: utf8mb4
-- Collation: utf8mb4_unicode_ci

USE sec_baseball;

-- ============================================================
-- Table: teams
-- Purpose: Stores information about all 16 SEC baseball teams
-- ============================================================

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

-- ============================================================
-- Table: players
-- Purpose: Stores information about all SEC baseball players
-- ============================================================

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

-- ============================================================
-- Table: games
-- Purpose: Stores information about SEC baseball games
-- ============================================================

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

-- ============================================================
-- Table: standings
-- Purpose: Stores calculated conference standings
-- ============================================================

CREATE TABLE standings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    team_id INT NOT NULL UNIQUE,
    standing_rank INT NOT NULL,
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

CREATE INDEX idx_standing_rank ON standings(standing_rank);
