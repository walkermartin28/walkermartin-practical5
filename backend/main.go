package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// enableCORS adds CORS headers to the response
func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "3600")
}

// corsMiddleware wraps handlers with CORS support
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w, r)

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// healthHandler handles GET /health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "SEC Baseball Hub API running"}`))
}

// apiRouter routes API requests to appropriate handlers
func apiRouter(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	path := r.URL.Path
	method := r.Method

	// Teams endpoints
	if strings.HasPrefix(path, "/api/teams") {
		if path == "/api/teams" {
			switch method {
			case "GET":
				GetTeams(w, r)
			case "POST":
				CreateTeam(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		} else if strings.HasPrefix(path, "/api/teams/") {
			switch method {
			case "GET":
				GetTeamByID(w, r)
			case "PUT":
				UpdateTeam(w, r)
			case "DELETE":
				DeleteTeam(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		}
		return
	}

	// Players endpoints
	if strings.HasPrefix(path, "/api/players") {
		if path == "/api/players" {
			switch method {
			case "GET":
				GetPlayers(w, r)
			case "POST":
				CreatePlayer(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		} else if strings.HasPrefix(path, "/api/players/") {
			switch method {
			case "GET":
				GetPlayerByID(w, r)
			case "PUT":
				UpdatePlayer(w, r)
			case "DELETE":
				DeletePlayer(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		}
		return
	}

	// Games endpoints
	if strings.HasPrefix(path, "/api/games") {
		if path == "/api/games" {
			switch method {
			case "GET":
				GetGames(w, r)
			case "POST":
				CreateGame(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		} else if strings.HasPrefix(path, "/api/games/") {
			switch method {
			case "GET":
				GetGameByID(w, r)
			case "PUT":
				UpdateGame(w, r)
			case "DELETE":
				DeleteGame(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		}
		return
	}

	// Standings endpoints
	if strings.HasPrefix(path, "/api/standings") {
		if path == "/api/standings" {
			switch method {
			case "GET":
				GetStandings(w, r)
			case "POST":
				CreateStanding(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		} else if strings.HasPrefix(path, "/api/standings/") {
			switch method {
			case "GET":
				GetStandingByID(w, r)
			case "PUT":
				UpdateStanding(w, r)
			case "DELETE":
				DeleteStanding(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		}
		return
	}

	// Auth endpoints
	if strings.HasPrefix(path, "/api/auth") {
		if path == "/api/auth/login" {
			if method == "POST" {
				Login(w, r)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return
		} else if path == "/api/auth/verify" {
			if method == "GET" {
				VerifyToken(w, r)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}
	}

	// Not found
	http.Error(w, "Endpoint not found", http.StatusNotFound)
}

func main() {
	// Initialize database connection
	log.Println("Connecting to database...")
	if err := InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer CloseDB()

	// Set up routes
	http.HandleFunc("/health", corsMiddleware(healthHandler))
	http.HandleFunc("/api/", apiRouter)

	// Start server
	port := "8080"
	log.Printf("SEC Baseball Hub API server starting on port %s...", port)
	log.Printf("Health check available at http://localhost:%s/health", port)
	log.Printf("API endpoints available at http://localhost:%s/api/", port)

	// Start HTTP server in a goroutine
	server := &http.Server{
		Addr: ":" + port,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	log.Println("Server is running. Press Ctrl+C to stop.")

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}
