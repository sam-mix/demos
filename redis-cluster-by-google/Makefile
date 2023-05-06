.PHONY: gen-conf
gen-conf:
	@./0-gen-conf.sh

.PHONY: up
up:
	@docker-compose up -d
	@echo 进入十秒休眠
	@sleep 10
	@echo 创建集群中....
	@docker exec -it redis-6371 redis-cli -h 192.168.0.219 -p 6371 -a 123456 --cluster create 192.168.0.219:6371 192.168.0.219:6372 192.168.0.219:6373 192.168.0.219:6374 192.168.0.219:6375 192.168.0.219:6376 --cluster-replicas 1
	@echo 创建集群完成
	@echo redis-cli hello =================================================
	@docker exec -it redis-6371 redis-cli -h 192.168.0.219 -p 6371 -a 123456 hello
	@echo 查看集群结点信息 ==================================================
	@docker exec -it redis-6371 redis-cli -h 192.168.0.219 -p 6371 -a 123456 cluster nodes
	@echo 查看集群信息 ======================================================
	@docker exec -it redis-6371 redis-cli -h 192.168.0.219 -p 6371 -a 123456 cluster info

.PHONY: down
down:
	@docker-compose down

.PHONY: test
test:
	@sleep 5 
	@echo hello


