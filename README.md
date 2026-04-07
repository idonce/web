# idonce Web

Landing page and documentation website for the idonce human verification system.

## What it does

- Serves the **landing page** at `/`
- **Developer guide** at `/developers` with integration docs, use cases, API reference
- **Legal pages**: `/privacy`, `/terms`, `/impressum`
- **Style guide** for web and mobile design reference
- Self-hosted fonts (Inter, Space Grotesk)

## Endpoints

| Method | Path | Description |
|---|---|---|
| `GET` | `/` | Landing page |
| `GET` | `/developers` | Developer documentation & integration guide |
| `GET` | `/privacy` | Privacy policy |
| `GET` | `/terms` | Terms of service |
| `GET` | `/impressum` | Legal notice |
| `GET` | `/fonts/*` | Self-hosted fonts |
| `GET` | `/health` | Health check |

## Build

```bash
# Generate static files to dist/
go run . --build

# Output:
# dist/index.html
# dist/developers.html
# dist/privacy.html
# dist/terms.html
# dist/impressum.html
# dist/fonts/
# dist/styleguide.html
```

Host `dist/` with any static file server (Caddy, nginx, GitHub Pages, Netlify, S3).

## Dev Server

```bash
go run .
# Listening on :9091
# Open http://localhost:9091
```

## Test

```bash
go test ./... -v
go vet ./...
staticcheck ./...
```

## Deploy

```bash
# Docker (uses Caddy to serve static files)
docker build -t idonce-web .
docker run -p 9091:9091 idonce-web

# Or just host dist/ directly
go run . --build
caddy file-server --root dist --listen :9091
```

## Related

- [idonce-issuer](https://github.com/idonce/issuer) — Credential Issuer (OpenID4VCI)
- [idonce-verifier](https://github.com/idonce/verifier) — Credential Verifier (OpenID4VP)

## Tech

- Go 1.22, zero external dependencies
- HTML/CSS embedded as Go constants
- Self-hosted fonts (Inter, Space Grotesk)
- Sidebar navigation with scroll-spy
