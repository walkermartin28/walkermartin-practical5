package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ErrorResponse represents a JSON error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// SuccessResponse represents a generic success response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Count   int         `json:"count,omitempty"`
}

// Team represents a SEC baseball team
type Team struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Abbreviation     string `json:"abbreviation"`
	LogoURL          string `json:"logo_url,omitempty"`
	HomeStadium      string `json:"home_stadium,omitempty"`
	Location         string `json:"location,omitempty"`
	ConferenceWins   int    `json:"conference_wins"`
	ConferenceLosses int    `json:"conference_losses"`
	OverallWins      int    `json:"overall_wins"`
	OverallLosses    int    `json:"overall_losses"`
	RunsScored       int    `json:"runs_scored"`
	RunsAllowed      int    `json:"runs_allowed"`
}

// Player represents a SEC baseball player
type Player struct {
	ID             int     `json:"id"`
	TeamID         int     `json:"team_id"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	JerseyNumber   *int    `json:"jersey_number,omitempty"`
	Position       string  `json:"position"`
	Year           string  `json:"year,omitempty"`
	Height         string  `json:"height,omitempty"`
	Weight         *int    `json:"weight,omitempty"`
	Hometown       string  `json:"hometown,omitempty"`
	HighSchool     string  `json:"high_school,omitempty"`
	BattingAvg     *float64 `json:"batting_avg,omitempty"`
	HomeRuns       int     `json:"home_runs"`
	RBIs           int     `json:"rbis"`
	StolenBases    int     `json:"stolen_bases"`
	ERA            *float64 `json:"era,omitempty"`
	Strikeouts     int     `json:"strikeouts"`
	Wins           int     `json:"wins"`
	Saves          int     `json:"saves"`
	InningsPitched *float64 `json:"innings_pitched,omitempty"`
	GamesPlayed    int     `json:"games_played"`
}

// Game represents a SEC baseball game
type Game struct {
	ID               int     `json:"id"`
	HomeTeamID       int     `json:"home_team_id"`
	AwayTeamID       int     `json:"away_team_id"`
	GameDate         string  `json:"game_date"`
	GameTime         *string `json:"game_time,omitempty"`
	Location         string  `json:"location,omitempty"`
	Status           string  `json:"status"`
	HomeScore        *int    `json:"home_score,omitempty"`
	AwayScore        *int    `json:"away_score,omitempty"`
	Innings          int     `json:"innings"`
	IsConferenceGame bool    `json:"is_conference_game"`
	Attendance       *int    `json:"attendance,omitempty"`
	Notes            string  `json:"notes,omitempty"`
}

// Standing represents a team's conference standing
type Standing struct {
	ID                    int      `json:"id"`
	TeamID                int      `json:"team_id"`
	StandingRank          int      `json:"standing_rank"`
	ConferenceWinPct      *float64 `json:"conference_win_pct,omitempty"`
	OverallWinPct         *float64 `json:"overall_win_pct,omitempty"`
	HomeRecord            string   `json:"home_record,omitempty"`
	AwayRecord            string   `json:"away_record,omitempty"`
	NeutralRecord         string   `json:"neutral_record,omitempty"`
	Streak                string   `json:"streak,omitempty"`
	Last10                string   `json:"last_10,omitempty"`
	GamesBack             float64  `json:"games_back"`
	RunsPerGame           *float64 `json:"runs_per_game,omitempty"`
	RunsAllowedPerGame    *float64 `json:"runs_allowed_per_game,omitempty"`
}

// sendJSON sends a JSON response
func sendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// sendError sends an error response
func sendError(w http.ResponseWriter, message string, statusCode int) {
	sendJSON(w, ErrorResponse{Success: false, Error: message}, statusCode)
}

// getIDFromPath extracts the ID from the URL path
func getIDFromPath(path string) (int, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 3 {
		return 0, nil
	}
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, err
	}
	return id, nil
}

// ========================================
// TEAMS HANDLERS
// ========================================

// GetTeams handles GET /api/teams
func GetTeams(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		SELECT id, name, abbreviation, logo_url, home_stadium, location,
		       conference_wins, conference_losses, overall_wins, overall_losses,
		       runs_scored, runs_allowed
		FROM teams
		ORDER BY name
	`)
	if err != nil {
		log.Printf("Error querying teams: %v", err)
		sendError(w, "Failed to retrieve teams", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	teams := []Team{}
	for rows.Next() {
		var t Team
		var logoURL, homeStadium, location sql.NullString

		err := rows.Scan(&t.ID, &t.Name, &t.Abbreviation, &logoURL, &homeStadium, &location,
			&t.ConferenceWins, &t.ConferenceLosses, &t.OverallWins, &t.OverallLosses,
			&t.RunsScored, &t.RunsAllowed)
		if err != nil {
			log.Printf("Error scanning team: %v", err)
			continue
		}

		if logoURL.Valid {
			t.LogoURL = logoURL.String
		}
		if homeStadium.Valid {
			t.HomeStadium = homeStadium.String
		}
		if location.Valid {
			t.Location = location.String
		}

		teams = append(teams, t)
	}

	sendJSON(w, SuccessResponse{Success: true, Data: teams, Count: len(teams)}, http.StatusOK)
}

// GetTeamByID handles GET /api/teams/{id}
func GetTeamByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	var t Team
	var logoURL, homeStadium, location sql.NullString

	err = db.QueryRow(`
		SELECT id, name, abbreviation, logo_url, home_stadium, location,
		       conference_wins, conference_losses, overall_wins, overall_losses,
		       runs_scored, runs_allowed
		FROM teams
		WHERE id = ?
	`, id).Scan(&t.ID, &t.Name, &t.Abbreviation, &logoURL, &homeStadium, &location,
		&t.ConferenceWins, &t.ConferenceLosses, &t.OverallWins, &t.OverallLosses,
		&t.RunsScored, &t.RunsAllowed)

	if err == sql.ErrNoRows {
		sendError(w, "Team not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Error querying team: %v", err)
		sendError(w, "Failed to retrieve team", http.StatusInternalServerError)
		return
	}

	if logoURL.Valid {
		t.LogoURL = logoURL.String
	}
	if homeStadium.Valid {
		t.HomeStadium = homeStadium.String
	}
	if location.Valid {
		t.Location = location.String
	}

	sendJSON(w, SuccessResponse{Success: true, Data: t}, http.StatusOK)
}

// CreateTeam handles POST /api/teams
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var t Team
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`
		INSERT INTO teams (name, abbreviation, logo_url, home_stadium, location,
		                   conference_wins, conference_losses, overall_wins, overall_losses,
		                   runs_scored, runs_allowed)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, t.Name, t.Abbreviation, t.LogoURL, t.HomeStadium, t.Location,
		t.ConferenceWins, t.ConferenceLosses, t.OverallWins, t.OverallLosses,
		t.RunsScored, t.RunsAllowed)

	if err != nil {
		log.Printf("Error creating team: %v", err)
		sendError(w, "Failed to create team", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	t.ID = int(id)

	sendJSON(w, SuccessResponse{Success: true, Data: t}, http.StatusCreated)
}

// UpdateTeam handles PUT /api/teams/{id}
func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	var t Team
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`
		UPDATE teams
		SET name = ?, abbreviation = ?, logo_url = ?, home_stadium = ?, location = ?,
		    conference_wins = ?, conference_losses = ?, overall_wins = ?, overall_losses = ?,
		    runs_scored = ?, runs_allowed = ?
		WHERE id = ?
	`, t.Name, t.Abbreviation, t.LogoURL, t.HomeStadium, t.Location,
		t.ConferenceWins, t.ConferenceLosses, t.OverallWins, t.OverallLosses,
		t.RunsScored, t.RunsAllowed, id)

	if err != nil {
		log.Printf("Error updating team: %v", err)
		sendError(w, "Failed to update team", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sendError(w, "Team not found", http.StatusNotFound)
		return
	}

	t.ID = id
	sendJSON(w, SuccessResponse{Success: true, Data: t}, http.StatusOK)
}

// DeleteTeam handles DELETE /api/teams/{id}
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("DELETE FROM teams WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting team: %v", err)
		sendError(w, "Failed to delete team", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sendError(w, "Team not found", http.StatusNotFound)
		return
	}

	sendJSON(w, SuccessResponse{Success: true, Data: map[string]string{"message": "Team deleted successfully"}}, http.StatusOK)
}

// ========================================
// PLAYERS HANDLERS
// ========================================

// GetPlayers handles GET /api/players
func GetPlayers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		SELECT id, team_id, first_name, last_name, jersey_number, position, year,
		       height, weight, hometown, high_school, batting_avg, home_runs, rbis,
		       stolen_bases, era, strikeouts, wins, saves, innings_pitched, games_played
		FROM players
		ORDER BY last_name, first_name
	`)
	if err != nil {
		log.Printf("Error querying players: %v", err)
		sendError(w, "Failed to retrieve players", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	players := []Player{}
	for rows.Next() {
		var p Player
		var jerseyNumber, weight sql.NullInt64
		var year, height, hometown, highSchool sql.NullString
		var battingAvg, era, inningsPitched sql.NullFloat64

		err := rows.Scan(&p.ID, &p.TeamID, &p.FirstName, &p.LastName, &jerseyNumber, &p.Position, &year,
			&height, &weight, &hometown, &highSchool, &battingAvg, &p.HomeRuns, &p.RBIs,
			&p.StolenBases, &era, &p.Strikeouts, &p.Wins, &p.Saves, &inningsPitched, &p.GamesPlayed)
		if err != nil {
			log.Printf("Error scanning player: %v", err)
			continue
		}

		if jerseyNumber.Valid {
			num := int(jerseyNumber.Int64)
			p.JerseyNumber = &num
		}
		if weight.Valid {
			w := int(weight.Int64)
			p.Weight = &w
		}
		if year.Valid {
			p.Year = year.String
		}
		if height.Valid {
			p.Height = height.String
		}
		if hometown.Valid {
			p.Hometown = hometown.String
		}
		if highSchool.Valid {
			p.HighSchool = highSchool.String
		}
		if battingAvg.Valid {
			p.BattingAvg = &battingAvg.Float64
		}
		if era.Valid {
			p.ERA = &era.Float64
		}
		if inningsPitched.Valid {
			p.InningsPitched = &inningsPitched.Float64
		}

		players = append(players, p)
	}

	sendJSON(w, SuccessResponse{Success: true, Data: players, Count: len(players)}, http.StatusOK)
}

// GetPlayerByID handles GET /api/players/{id}
func GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	var p Player
	var jerseyNumber, weight sql.NullInt64
	var year, height, hometown, highSchool sql.NullString
	var battingAvg, era, inningsPitched sql.NullFloat64

	err = db.QueryRow(`
		SELECT id, team_id, first_name, last_name, jersey_number, position, year,
		       height, weight, hometown, high_school, batting_avg, home_runs, rbis,
		       stolen_bases, era, strikeouts, wins, saves, innings_pitched, games_played
		FROM players
		WHERE id = ?
	`, id).Scan(&p.ID, &p.TeamID, &p.FirstName, &p.LastName, &jerseyNumber, &p.Position, &year,
		&height, &weight, &hometown, &highSchool, &battingAvg, &p.HomeRuns, &p.RBIs,
		&p.StolenBases, &era, &p.Strikeouts, &p.Wins, &p.Saves, &inningsPitched, &p.GamesPlayed)

	if err == sql.ErrNoRows {
		sendError(w, "Player not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Error querying player: %v", err)
		sendError(w, "Failed to retrieve player", http.StatusInternalServerError)
		return
	}

	if jerseyNumber.Valid {
		num := int(jerseyNumber.Int64)
		p.JerseyNumber = &num
	}
	if weight.Valid {
		w := int(weight.Int64)
		p.Weight = &w
	}
	if year.Valid {
		p.Year = year.String
	}
	if height.Valid {
		p.Height = height.String
	}
	if hometown.Valid {
		p.Hometown = hometown.String
	}
	if highSchool.Valid {
		p.HighSchool = highSchool.String
	}
	if battingAvg.Valid {
		p.BattingAvg = &battingAvg.Float64
	}
	if era.Valid {
		p.ERA = &era.Float64
	}
	if inningsPitched.Valid {
		p.InningsPitched = &inningsPitched.Float64
	}

	sendJSON(w, SuccessResponse{Success: true, Data: p}, http.StatusOK)
}

// CreatePlayer handles POST /api/players
func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var p Player
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`
		INSERT INTO players (team_id, first_name, last_name, jersey_number, position, year,
		                     height, weight, hometown, high_school, batting_avg, home_runs, rbis,
		                     stolen_bases, era, strikeouts, wins, saves, innings_pitched, games_played)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, p.TeamID, p.FirstName, p.LastName, p.JerseyNumber, p.Position, p.Year,
		p.Height, p.Weight, p.Hometown, p.HighSchool, p.BattingAvg, p.HomeRuns, p.RBIs,
		p.StolenBases, p.ERA, p.Strikeouts, p.Wins, p.Saves, p.InningsPitched, p.GamesPlayed)

	if err != nil {
		log.Printf("Error creating player: %v", err)
		sendError(w, "Failed to create player", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	p.ID = int(id)

	sendJSON(w, SuccessResponse{Success: true, Data: p}, http.StatusCreated)
}

// UpdatePlayer handles PUT /api/players/{id}
func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	var p Player
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`
		UPDATE players
		SET team_id = ?, first_name = ?, last_name = ?, jersey_number = ?, position = ?, year = ?,
		    height = ?, weight = ?, hometown = ?, high_school = ?, batting_avg = ?, home_runs = ?, rbis = ?,
		    stolen_bases = ?, era = ?, strikeouts = ?, wins = ?, saves = ?, innings_pitched = ?, games_played = ?
		WHERE id = ?
	`, p.TeamID, p.FirstName, p.LastName, p.JerseyNumber, p.Position, p.Year,
		p.Height, p.Weight, p.Hometown, p.HighSchool, p.BattingAvg, p.HomeRuns, p.RBIs,
		p.StolenBases, p.ERA, p.Strikeouts, p.Wins, p.Saves, p.InningsPitched, p.GamesPlayed, id)

	if err != nil {
		log.Printf("Error updating player: %v", err)
		sendError(w, "Failed to update player", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sendError(w, "Player not found", http.StatusNotFound)
		return
	}

	p.ID = id
	sendJSON(w, SuccessResponse{Success: true, Data: p}, http.StatusOK)
}

// DeletePlayer handles DELETE /api/players/{id}
func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("DELETE FROM players WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting player: %v", err)
		sendError(w, "Failed to delete player", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sendError(w, "Player not found", http.StatusNotFound)
		return
	}

	sendJSON(w, SuccessResponse{Success: true, Data: map[string]string{"message": "Player deleted successfully"}}, http.StatusOK)
}

// ========================================
// GAMES HANDLERS
// ========================================

// GetGames handles GET /api/games
func GetGames(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		SELECT id, home_team_id, away_team_id, game_date, game_time, location,
		       status, home_score, away_score, innings, is_conference_game, attendance, notes
		FROM games
		ORDER BY game_date DESC
	`)
	if err != nil {
		log.Printf("Error querying games: %v", err)
		sendError(w, "Failed to retrieve games", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	games := []Game{}
	for rows.Next() {
		var g Game
		var gameTime, location, notes sql.NullString
		var homeScore, awayScore, attendance sql.NullInt64

		err := rows.Scan(&g.ID, &g.HomeTeamID, &g.AwayTeamID, &g.GameDate, &gameTime, &location,
			&g.Status, &homeScore, &awayScore, &g.Innings, &g.IsConferenceGame, &attendance, &notes)
		if err != nil {
			log.Printf("Error scanning game: %v", err)
			continue
		}

		if gameTime.Valid {
			g.GameTime = &gameTime.String
		}
		if location.Valid {
			g.Location = location.String
		}
		if homeScore.Valid {
			score := int(homeScore.Int64)
			g.HomeScore = &score
		}
		if awayScore.Valid {
			score := int(awayScore.Int64)
			g.AwayScore = &score
		}
		if attendance.Valid {
			att := int(attendance.Int64)
			g.Attendance = &att
		}
		if notes.Valid {
			g.Notes = notes.String
		}

		games = append(games, g)
	}

	sendJSON(w, SuccessResponse{Success: true, Data: games, Count: len(games)}, http.StatusOK)
}

// GetGameByID handles GET /api/games/{id}
func GetGameByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	var g Game
	var gameTime, location, notes sql.NullString
	var homeScore, awayScore, attendance sql.NullInt64

	err = db.QueryRow(`
		SELECT id, home_team_id, away_team_id, game_date, game_time, location,
		       status, home_score, away_score, innings, is_conference_game, attendance, notes
		FROM games
		WHERE id = ?
	`, id).Scan(&g.ID, &g.HomeTeamID, &g.AwayTeamID, &g.GameDate, &gameTime, &location,
		&g.Status, &homeScore, &awayScore, &g.Innings, &g.IsConferenceGame, &attendance, &notes)

	if err == sql.ErrNoRows {
		sendError(w, "Game not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Error querying game: %v", err)
		sendError(w, "Failed to retrieve game", http.StatusInternalServerError)
		return
	}

	if gameTime.Valid {
		g.GameTime = &gameTime.String
	}
	if location.Valid {
		g.Location = location.String
	}
	if homeScore.Valid {
		score := int(homeScore.Int64)
		g.HomeScore = &score
	}
	if awayScore.Valid {
		score := int(awayScore.Int64)
		g.AwayScore = &score
	}
	if attendance.Valid {
		att := int(attendance.Int64)
		g.Attendance = &att
	}
	if notes.Valid {
		g.Notes = notes.String
	}

	sendJSON(w, SuccessResponse{Success: true, Data: g}, http.StatusOK)
}

// CreateGame handles POST /api/games
func CreateGame(w http.ResponseWriter, r *http.Request) {
	var g Game
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`
		INSERT INTO games (home_team_id, away_team_id, game_date, game_time, location,
		                   status, home_score, away_score, innings, is_conference_game, attendance, notes)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, g.HomeTeamID, g.AwayTeamID, g.GameDate, g.GameTime, g.Location,
		g.Status, g.HomeScore, g.AwayScore, g.Innings, g.IsConferenceGame, g.Attendance, g.Notes)

	if err != nil {
		log.Printf("Error creating game: %v", err)
		sendError(w, "Failed to create game", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	g.ID = int(id)

	sendJSON(w, SuccessResponse{Success: true, Data: g}, http.StatusCreated)
}

// UpdateGame handles PUT /api/games/{id}
func UpdateGame(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	var g Game
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`
		UPDATE games
		SET home_team_id = ?, away_team_id = ?, game_date = ?, game_time = ?, location = ?,
		    status = ?, home_score = ?, away_score = ?, innings = ?, is_conference_game = ?, attendance = ?, notes = ?
		WHERE id = ?
	`, g.HomeTeamID, g.AwayTeamID, g.GameDate, g.GameTime, g.Location,
		g.Status, g.HomeScore, g.AwayScore, g.Innings, g.IsConferenceGame, g.Attendance, g.Notes, id)

	if err != nil {
		log.Printf("Error updating game: %v", err)
		sendError(w, "Failed to update game", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sendError(w, "Game not found", http.StatusNotFound)
		return
	}

	g.ID = id
	sendJSON(w, SuccessResponse{Success: true, Data: g}, http.StatusOK)
}

// DeleteGame handles DELETE /api/games/{id}
func DeleteGame(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("DELETE FROM games WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting game: %v", err)
		sendError(w, "Failed to delete game", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sendError(w, "Game not found", http.StatusNotFound)
		return
	}

	sendJSON(w, SuccessResponse{Success: true, Data: map[string]string{"message": "Game deleted successfully"}}, http.StatusOK)
}

// ========================================
// STANDINGS HANDLERS
// ========================================

// GetStandings handles GET /api/standings
func GetStandings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		SELECT id, team_id, standing_rank, conference_win_pct, overall_win_pct,
		       home_record, away_record, neutral_record, streak, last_10,
		       games_back, runs_per_game, runs_allowed_per_game
		FROM standings
		ORDER BY standing_rank
	`)
	if err != nil {
		log.Printf("Error querying standings: %v", err)
		sendError(w, "Failed to retrieve standings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	standings := []Standing{}
	for rows.Next() {
		var s Standing
		var conferenceWinPct, overallWinPct, runsPerGame, runsAllowedPerGame sql.NullFloat64
		var homeRecord, awayRecord, neutralRecord, streak, last10 sql.NullString

		err := rows.Scan(&s.ID, &s.TeamID, &s.StandingRank, &conferenceWinPct, &overallWinPct,
			&homeRecord, &awayRecord, &neutralRecord, &streak, &last10,
			&s.GamesBack, &runsPerGame, &runsAllowedPerGame)
		if err != nil {
			log.Printf("Error scanning standing: %v", err)
			continue
		}

		if conferenceWinPct.Valid {
			s.ConferenceWinPct = &conferenceWinPct.Float64
		}
		if overallWinPct.Valid {
			s.OverallWinPct = &overallWinPct.Float64
		}
		if homeRecord.Valid {
			s.HomeRecord = homeRecord.String
		}
		if awayRecord.Valid {
			s.AwayRecord = awayRecord.String
		}
		if neutralRecord.Valid {
			s.NeutralRecord = neutralRecord.String
		}
		if streak.Valid {
			s.Streak = streak.String
		}
		if last10.Valid {
			s.Last10 = last10.String
		}
		if runsPerGame.Valid {
			s.RunsPerGame = &runsPerGame.Float64
		}
		if runsAllowedPerGame.Valid {
			s.RunsAllowedPerGame = &runsAllowedPerGame.Float64
		}

		standings = append(standings, s)
	}

	sendJSON(w, SuccessResponse{Success: true, Data: standings, Count: len(standings)}, http.StatusOK)
}

// GetStandingByID handles GET /api/standings/{id}
func GetStandingByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid standing ID", http.StatusBadRequest)
		return
	}

	var s Standing
	var conferenceWinPct, overallWinPct, runsPerGame, runsAllowedPerGame sql.NullFloat64
	var homeRecord, awayRecord, neutralRecord, streak, last10 sql.NullString

	err = db.QueryRow(`
		SELECT id, team_id, standing_rank, conference_win_pct, overall_win_pct,
		       home_record, away_record, neutral_record, streak, last_10,
		       games_back, runs_per_game, runs_allowed_per_game
		FROM standings
		WHERE id = ?
	`, id).Scan(&s.ID, &s.TeamID, &s.StandingRank, &conferenceWinPct, &overallWinPct,
		&homeRecord, &awayRecord, &neutralRecord, &streak, &last10,
		&s.GamesBack, &runsPerGame, &runsAllowedPerGame)

	if err == sql.ErrNoRows {
		sendError(w, "Standing not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Error querying standing: %v", err)
		sendError(w, "Failed to retrieve standing", http.StatusInternalServerError)
		return
	}

	if conferenceWinPct.Valid {
		s.ConferenceWinPct = &conferenceWinPct.Float64
	}
	if overallWinPct.Valid {
		s.OverallWinPct = &overallWinPct.Float64
	}
	if homeRecord.Valid {
		s.HomeRecord = homeRecord.String
	}
	if awayRecord.Valid {
		s.AwayRecord = awayRecord.String
	}
	if neutralRecord.Valid {
		s.NeutralRecord = neutralRecord.String
	}
	if streak.Valid {
		s.Streak = streak.String
	}
	if last10.Valid {
		s.Last10 = last10.String
	}
	if runsPerGame.Valid {
		s.RunsPerGame = &runsPerGame.Float64
	}
	if runsAllowedPerGame.Valid {
		s.RunsAllowedPerGame = &runsAllowedPerGame.Float64
	}

	sendJSON(w, SuccessResponse{Success: true, Data: s}, http.StatusOK)
}

// CreateStanding handles POST /api/standings
func CreateStanding(w http.ResponseWriter, r *http.Request) {
	var s Standing
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`
		INSERT INTO standings (team_id, standing_rank, conference_win_pct, overall_win_pct,
		                       home_record, away_record, neutral_record, streak, last_10,
		                       games_back, runs_per_game, runs_allowed_per_game)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, s.TeamID, s.StandingRank, s.ConferenceWinPct, s.OverallWinPct,
		s.HomeRecord, s.AwayRecord, s.NeutralRecord, s.Streak, s.Last10,
		s.GamesBack, s.RunsPerGame, s.RunsAllowedPerGame)

	if err != nil {
		log.Printf("Error creating standing: %v", err)
		sendError(w, "Failed to create standing", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	s.ID = int(id)

	sendJSON(w, SuccessResponse{Success: true, Data: s}, http.StatusCreated)
}

// UpdateStanding handles PUT /api/standings/{id}
func UpdateStanding(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid standing ID", http.StatusBadRequest)
		return
	}

	var s Standing
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`
		UPDATE standings
		SET team_id = ?, standing_rank = ?, conference_win_pct = ?, overall_win_pct = ?,
		    home_record = ?, away_record = ?, neutral_record = ?, streak = ?, last_10 = ?,
		    games_back = ?, runs_per_game = ?, runs_allowed_per_game = ?
		WHERE id = ?
	`, s.TeamID, s.StandingRank, s.ConferenceWinPct, s.OverallWinPct,
		s.HomeRecord, s.AwayRecord, s.NeutralRecord, s.Streak, s.Last10,
		s.GamesBack, s.RunsPerGame, s.RunsAllowedPerGame, id)

	if err != nil {
		log.Printf("Error updating standing: %v", err)
		sendError(w, "Failed to update standing", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sendError(w, "Standing not found", http.StatusNotFound)
		return
	}

	s.ID = id
	sendJSON(w, SuccessResponse{Success: true, Data: s}, http.StatusOK)
}

// DeleteStanding handles DELETE /api/standings/{id}
func DeleteStanding(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil || id == 0 {
		sendError(w, "Invalid standing ID", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("DELETE FROM standings WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting standing: %v", err)
		sendError(w, "Failed to delete standing", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sendError(w, "Standing not found", http.StatusNotFound)
		return
	}

	sendJSON(w, SuccessResponse{Success: true, Data: map[string]string{"message": "Standing deleted successfully"}}, http.StatusOK)
}
