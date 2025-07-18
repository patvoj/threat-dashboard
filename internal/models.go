package models

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
