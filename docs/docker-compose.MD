# DOCKER COMPOSE

docker-compose.yaml
- postgres
- adminer(http://localhost:8080)
docker-com

docker compose up -d
docker compose ps
docker compose exec db psql -U postgres -c '\l'
docker compose down