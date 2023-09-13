# INSTALL SCOOP WINDOWS

Set-ExecutionPolicy RemoteSigned -Scope CurrentUser # Optional: Needed to run a remote script the first time

irm get.scoop.sh | iex

scoop install migrate

# CREATE MIGRATIONS

migrate create -ext sql -dir db/migration -seq initial_tables
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/email_sender?sslmode=disable" -verbose up