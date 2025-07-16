package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) index(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello from the server")
		fmt.Println("File path is:", filePath)

		jsonData := `{
			"threatName": "Win32/Rbot",
			"category": "trojan",
			"size": 437289,
			"detectionDate": "2019-04-01",
			"variants": [
				{
					"name": "Win32/TrojanProxy.Emotet.A",
					"dateAdded": "2019-04-10"
				},
				{
					"name": "Win32/TrojanProxy.Emotet.B",
					"dateAdded": "2019-04-22"
				}
			]
		}`

		var threat ThreatData
		if err := json.Unmarshal([]byte(jsonData), &threat); err != nil {
			http.Error(w, "Error parsing JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		tmpl, err := template.ParseFiles(filePath)
		if err != nil {
			http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")

		if err := tmpl.Execute(w, threat); err != nil {
			http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (app *application) render(w http.ResponseWriter, r *http.Request) {
}

type ThreatData struct {
	ThreatName    string          `json:"threatName"`
	Category      string          `json:"category"`
	Size          int             `json:"size"`
	DetectionDate string          `json:"detectionDate"`
	Variants      []ThreatVariant `json:"variants"`
}

type ThreatVariant struct {
	Name      string `json:"name"`
	DateAdded string `json:"dateAdded"`
}
