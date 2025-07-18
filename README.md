# Threat Dashboard

A minimal web application for submitting and viewing threat data via a JSON input. Built with the Go standard library only — no external dependencies.

## Features
- Submit threat information in JSON format through a web form
- Automatically stores each threat entry in a local file (`threats.jsonl`)
- Displays all submitted threats on a dashboard
- Supports multiple variants per threat
- Accepts partial JSON (missing fields will default to zero values)
- Ignores unknown fields in JSON input

## Getting Started
### Installation & Running
```bash
# Clone and run directly
git clone git@github.com:patvoj/threat-dashboard.git
cd threat-dashboard

# Run with default settings
go run ./cmd/threat-dashboard

# Run with custom configuration
go run ./cmd/threat-dashboard -p :8080 -t ./ui/templates/custom.html

# View all available options
go run ./cmd/threat-dashboard -h
```

Visit `http://localhost:4000` in your browser to access the dashboard.

### Command Line Options
| Flag | Default | Description |
|------|---------|-------------|
| `-p` | `:4000` | Server port number |
| `-t` | `./ui/templates/threat.html.tmpl` | Path to HTML template file |


## Data Format
### JSON Schema
Threats are stored using the following JSON structure:
```json
{
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
}
```


## File storage
### JSONL Format
Threats are persisted in `internal/threats.jsonl` using JSON Lines format:
- Each line contains one complete JSON object
- Human-readable and easily parseable
- Supports concurrent read operations
- Can be manually edited if needed


## Usage Guide
### Submitting Threats
1. Access the Dashboard: Navigate to http://localhost:4000
2. Prepare JSON Data: Format your threat data according to the schema
3. Submit via Form:
    - Paste valid JSON into the textarea
    - Click submit button
    - Page redirects to dashboard with new entry displayed


## File Structure
```
threat-dashboard/
├── cmd/
│   └── threat-dashboard/      # Main app entry point
├── internal/                  # Internal packages and storage
├── ui/
│   ├── templates/             # HTML templates
│   └── static/                # CSS styles
└── README.md
```
