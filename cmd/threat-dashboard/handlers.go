package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func (app *application) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		threats, err := loadAllThreats()
		if err != nil {
			http.Error(w, "Failed to load threats", http.StatusInternalServerError)
			return
		}

		data := struct {
			Threats []ThreatData
		}{
			Threats: threats,
		}

		tmpl, err := template.ParseFiles(app.templPath)
		if err != nil {
			http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, data)
	}
}

func (app *application) render(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	rawJSON := r.FormValue("json_input")

	var threat ThreatData
	if err := json.Unmarshal([]byte(rawJSON), &threat); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := app.saveThreat(threat); err != nil {
		http.Error(w, "Failed to save threat: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
