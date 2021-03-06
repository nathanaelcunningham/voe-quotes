dev:
	docker compose up
dev-d:
	docker compose up -d
migrate:
	docker exec -it backend migrate -path migrations/ -database 'mysql://quotes:quotes@tcp(mariadb:3306)/quotes' up
migrate-down:
	docker exec -it backend migrate -path migrations/ -database 'mysql://quotes:quotes@tcp(mariadb:3306)/quotes' down
