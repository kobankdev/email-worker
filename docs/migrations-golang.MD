# INSTALL SCOOP WINDOWS

Set-ExecutionPolicy RemoteSigned -Scope CurrentUser # Optional: Needed to run a remote script the first time

irm get.scoop.sh | iex

scoop install migrate

# CREATE MIGRATIONS

migrate create -ext sql -dir db/migration -seq initial_tables

# Up db
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/email_sender?sslmode=disable" -verbose up

# Down db
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/email_sender?sslmode=disable" -verbose down

# Forçar a versão da migração 
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/email_sender?sslmode=disable" -verbose force 1

# Verificar o versão da migração
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/email_sender?sslmode=disable" -verbose version




