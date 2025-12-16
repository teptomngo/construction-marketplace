\# Local Development Setup (Phase 0)



This document defines the required local environment for Phase 0.

Do not proceed to later phases until this setup is verified.



\## Operating System

\- Windows (development)

\- Linux (deployment target on AWS)



\## Required Tools



\### Git

\- Used for version control

\- Verify:

&nbsp; git --version



\### Docker Desktop

\- Used to run PostgreSQL locally

\- Must support Docker Compose

\- Verify:

&nbsp; docker version

&nbsp; docker compose version



\### Go

\- Backend language

\- Version: 1.22.x (locked for Phase 0)

\- Verify:

&nbsp; go version



\### PostgreSQL Client (Optional)

\- Used for manual inspection/debugging

\- Not required for normal development



\### API Testing Tool

\- Postman

\- Used for validating API endpoints



\## Rules

\- PostgreSQL must always run via Docker locally

\- No services should be installed directly on the host unless documented



