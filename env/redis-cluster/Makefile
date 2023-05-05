.PHONY: up
up:
	@docker network create redis-cluster
	@docker-compose up -d

.PHONY: down
down:
	@docker-compose down
	@docker network rm redis-cluster
