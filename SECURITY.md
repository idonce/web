# Security Policy

## Reporting a Vulnerability

If you discover a security vulnerability in idonce, please report it responsibly.

**Email:** security@idonce.com

Please include:
- Description of the vulnerability
- Steps to reproduce
- Potential impact

We will acknowledge receipt within 48 hours and aim to provide a fix within 7 days for critical issues.

## Scope

This policy covers:
- The idonce web server (`idonce-web`)
- SD-JWT-VC verification logic
- OpenID4VP session management
- Any credential or cryptographic handling

## Out of Scope

- Denial of service via rate limiting (already mitigated)
- Issues in third-party dependencies (we have zero Go dependencies)
- Social engineering

## Disclosure

We follow coordinated disclosure. Please do not open public issues for security vulnerabilities.
