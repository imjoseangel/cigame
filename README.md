# CI Game

Generates a Table of Committers for a given Project.

## Requirements

- Go 1.14

## Run

`$ go run cmd/cigame.go`

or

`$ docker run -p 8000:8000 imjoseangel/cigame`

Open `http://localhost:8000?owner={owner}&repo={repo}`

### Environment variables

- `PORT` defines the port the server listens on (default 8000)
- `GITHUB_TOKEN` to get stats for private repos you'll need a [github access token](https://github.com/settings/tokens)
- `MAPPINGS` Path to a JSON file with the keys of email addresses to aliases
