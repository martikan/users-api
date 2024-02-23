# USERS-API

API for testing distributed logging in local environment.

## First steps for local env

Create .env file in root:
```text
DATABASE_URL=postgres://user:password@host:port/db_name
LOG_LEVEL=info # Set to debug for more detailed logs
```

Create local database by this script:
```bash
docker run --name users-db -d -e POSTGRES_USER=users-api -e POSTGRES_PASSWORD=aaa -p 5432:5432 postgres:16-bookworm
```

Or run:
```bash
make create-local-db
```