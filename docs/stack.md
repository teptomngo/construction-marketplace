\# Stack Choices (Locked for Phase 0–2)



\## Core

\- Database: PostgreSQL (target version: 16.x)

\- Migrations: golang-migrate

\- Logging: Zerolog

\- API Testing: Postman



\## Environments

\- Local development: Windows

\- Deployment: Linux (AWS)



\## Rules

\- PostgreSQL runs locally via Docker Compose.

\- All schema changes must be performed via migrations.

\- No manual database changes are allowed.



