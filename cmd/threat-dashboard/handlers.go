package main

import (
	"encoding/json"
	"net/http"

	models "github.com/patvoj/threat-dashboard/internal"
)

// index handles GET requests to the root path.
// It loads all threats from storage file and renders them using the HTML template.
// Returns HTTP 405 for non GET requests.
func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.logger.Warn("Method not allowed", "method", r.Method, "path", r.URL.Path)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	threats, err := loadAllThreats()
	if err != nil {
		app.logger.Error("Failed to load threats", "error", err)
		http.Error(w, "Failed to load threats", http.StatusInternalServerError)
		return
	}

	data := struct {
		Threats []models.ThreatData
	}{
		Threats: threats,
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	if err := app.templates.Execute(w, data); err != nil {
		app.logger.Error("Template execution failed", "error", err)
		http.Error(w, "Template rendering failed", http.StatusInternalServerError)
		return
	}
}

// render handles POST requests to add new threat data.
// It parses JSON from form data, validates it, saves to storage, and redirects to index.
// Returns HTTP 405 for non POST requests.
func (app *application) render(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.logger.Error("Method not allowed", "method", r.Method, "path", r.URL.Path)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		app.logger.Error("Form parsing failed", "error", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	rawJSON := r.FormValue("json_input")
	if rawJSON == "" {
		app.logger.Error("Empty JSON input received")
		http.Error(w, "JSON input is required", http.StatusBadRequest)
		return
	}

	var threat models.ThreatData
	if err := json.Unmarshal([]byte(rawJSON), &threat); err != nil {
		app.logger.Error("JSON parsing failed", "error", err, "json", rawJSON)
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := app.saveThreat(threat); err != nil {
		app.logger.Error("Failed to save threat", "error", err, "threat", threat.ThreatName)
		http.Error(w, "Failed to save threat: "+err.Error(), http.StatusInternalServerError)
		return
	}

	app.logger.Info("Threat saved successfully", "threat", threat.ThreatName)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
