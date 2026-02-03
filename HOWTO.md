# HOWTO

## Prerequisites

- Go 1.24+

## Project Structure

```
cli/          CLI tool (Cobra-based)
server/       HTTP server exposing the same logic as a JSON API
logic/        Shared business logic used by both CLI and server
Makefile      Build/run commands
```

## CLI

### Build

```bash
make build
```

Produces a `ko` binary in the project root.

### Run a command directly (without building)

```bash
go run ./cli recycling
```

### Install locally

```bash
make install
```

Copies the built binary to `/usr/local/bin/ko/ko`. After installing, run commands with:

```bash
ko recycling
```

### Full publish (clean + build + install)

```bash
make publish
```

## Server

### Run locally

```bash
make server
```

Starts the server on `http://localhost:8080`.

### Endpoints

| Method | Path          | Description                          |
|--------|---------------|--------------------------------------|
| GET    | `/recycling`  | Returns the next recycling day as JSON |

### Example request

```bash
curl http://localhost:8080/recycling
```

Example response:

```json
{"type": "paper", "in_days": 3}
```

If no recycling days remain in the year:

```json
{"message": "No more recycling days this year."}
```

## Makefile Reference

| Target    | Description                              |
|-----------|------------------------------------------|
| `build`   | Build the CLI binary                     |
| `install` | Build and copy binary to install dir     |
| `publish` | Clean, build, and install                |
| `clean`   | Remove the built binary                  |
| `server`  | Run the HTTP server locally on `:8080`   |
