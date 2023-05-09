# Docker Compose 文件
DOCKER_COMPOSE_FILE = docker-compose.yaml

# 编译 Etcd 镜像
build:
    docker build -t quay.io/coreos/etcd:v3.4.0 .

# 启动 Etcd 集群
up:
    docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# 停止 Etcd 集群
down:
    docker-compose -f $(DOCKER_COMPOSE_FILE) down

# 显示 Etcd 集群运行状态
ps:
    docker-compose -f $(DOCKER_COMPOSE_FILE) ps

# 删除 Etcd 集群容器
clean:
    docker-compose -f $(DOCKER_COMPOSE_FILE) down --volumes