package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync"
	// time library added to handle event date validation
	"time"
)

// ---------------------------------------------------------------------------
// Domain
// ---------------------------------------------------------------------------

// LAB7A - Update user struct in main.go
// Needed for new fields and their validations
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// type EventDate
	EventDate string `json:"event_date"`
	Tickets   int    `json:"tickets"`
}

// ---------------------------------------------------------------------------
// Application
// ---------------------------------------------------------------------------

type application struct {
	logger *slog.Logger
	mu     sync.Mutex
	users  []User
	nextID int
}

// ---------------------------------------------------------------------------
// Handlers
// ---------------------------------------------------------------------------

// LAB7A - Update createUser handler in main.go
// Also needed as user structu has more fields to input
func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		EventDate string `json:"event_date"`
		Tickets   int    `json:"tickets"`
		// Agreed bool used to indicate if the user has agreed to the terms
		Agreed bool   `json:"agreed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input);
	err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Validation for 3 new fields based on constraints given
	// Field 1: Event must be future EventDate
	// Field 2: Tickets must be between 1 and 5
	// Field 3: Terms and conditions must be checked to proceed

	eventTime, err := time.Parse("2006-01-02", input.EventDate)
	if err != nil {
		http.Error(w, "invalid event_date format", http.StatusBadRequest)
		return
	}
	if !eventTime.After(time.Now()) {
		http.Error(w, "event_date must be in the future", http.StatusBadRequest)
		return
	}
	if input.Tickets < 1 || input.Tickets > 5 {
		http.Error(w, "tickets must be between 1 and 5", http.StatusBadRequest)
		return
	}
	if !input.Agreed {
		http.Error(w, "terms and conditions must be agreed to", http.StatusBadRequest)
		return
	}


	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.ToLower(strings.TrimSpace(input.Email))

	if input.Name == "" || input.Email == "" {
		http.Error(w, "name and email are required", http.StatusBadRequest)
		return
	}

	app.mu.Lock()
	app.nextID++
	user := User{
		ID:        app.nextID,
		Name:      input.Name,
		Email:     input.Email,
		EventDate: input.EventDate,
		Tickets:   input.Tickets,
	}
	app.users = append(app.users, user)
	app.mu.Unlock()

	app.logger.Info("user created", "id", user.ID, "name", user.Name, "email", user.Email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (app *application) listUsers(w http.ResponseWriter, r *http.Request) {
	app.mu.Lock()
	users := make([]User, len(app.users))
	copy(users, app.users)
	app.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
		users:  []User{},
		nextID: 0,
	}

	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("POST /api/users", app.createUser)
	mux.HandleFunc("GET /api/users", app.listUsers)

	// Static files — serves index.html, CSS, JS from ./static
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	addr := ":4000"
	logger.Info("starting server", "addr", addr)

	err := http.ListenAndServe(addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
