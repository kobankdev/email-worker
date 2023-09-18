# Pull da imagem Docker
docker pull sqlc/sqlc

# Criar
/db/sqlc
/db/query/

# Gerar CRUD
docker run --rm -v "${PWD}:/src" -w /src sqlc/sqlc generate