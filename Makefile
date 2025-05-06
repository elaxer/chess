include .env

migrate-up:
	migrate -path migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up
migrate-down:
	migrate -path migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" down