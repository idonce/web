# Contributing to idonce

We welcome contributions. Here's how to get started.

## Development Setup

```bash
git clone https://github.com/idonce/web
cd web
go run .
# Open http://localhost:9091
```

## Before Submitting

```bash
go build ./...
go test ./... -v
go vet ./...
staticcheck ./...
```

All checks must pass. Zero warnings.

## Pull Requests

- Keep PRs focused on a single change
- Include tests for new functionality
- Follow existing code style (no external dependencies, stdlib only)
- Update README.md if adding endpoints or changing behavior

## Code Style

- Go 1.22, standard library only — no external dependencies
- HTML/CSS embedded as Go constants
- ES256 (P-256) for all cryptographic operations
- Error messages lowercase, no trailing punctuation

## Issues

Use GitHub Issues for bug reports and feature requests. Include steps to reproduce for bugs.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
