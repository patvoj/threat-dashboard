package main

import (
	"bufio"
	"encoding/json"
	"os"
)

func (app *application) saveThreat(threat ThreatData) error {
	f, err := os.OpenFile("threats.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(threat)
}

func loadAllThreats() ([]ThreatData, error) {
	file, err := os.Open("threats.jsonl")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var threats []ThreatData
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var t ThreatData
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
