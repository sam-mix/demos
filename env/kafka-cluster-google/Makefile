.PHONY: up down clean

clean:
	@echo clean

down:
	@docker-compose down
	@docker network rm kafka-cluster

up:
	@docker network create --driver=bridge --subnet=172.28.0.0/16 --gateway=172.28.0.1 kafka-cluster
	@docker-compose up -d


