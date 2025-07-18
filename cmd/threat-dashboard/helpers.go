package main

import (
	"bufio"
	"encoding/json"
	"os"

	models "github.com/patvoj/threat-dashboard/internal"
)

// saveThreat appends a threat record to the JSONL data file.
// It creates the file if it doesn't exist and ensures proper JSON formatting.
// Returns an error if file operations fail.
func (app *application) saveThreat(threat models.ThreatData) error {
	f, err := os.OpenFile("internal/threats.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(threat)
}

// loadAllThreats reads all threat records from the JSONL data file.
// It parses each line as a separate JSON object and returns a slice of ThreatData.
// Returns an empty slice if the file doesn't exist.
func loadAllThreats() ([]models.ThreatData, error) {
	file, err := os.Open("internal/threats.jsonl")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var threats []models.ThreatData
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var t models.ThreatData
		if err := json.Unmarshal(scanner.Bytes(), &t); err != nil {
			continue
		}
		threats = append(threats, t)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return threats, nil
}
