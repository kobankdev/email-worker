# RODAR CONTAINERS
docker compose up -d

# EXECUTAR MIGRATIONS
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/email_sender?sslmode=disable" -verbose up 

# ERRO MIGRATIONS
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/email_sender?sslmode=disable" -verbose force 1

# EXCLUIR DIRETÓRIO FORÇA WINDOWS
Remove-Item -Path .\data -Recurse -Force

# VERIFICAR VERSÃO MIGRATIONS
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/email_sender?sslmode=disable" -verbose version