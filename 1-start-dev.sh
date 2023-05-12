#!/usr/bin/env bash

set -e

# 配置文件路径
CONFIG_ROOT_PATH=./conf
# 配置文件应用到的环境
CONFIG_APPLY_ENV=dev
# Redis配置文件名称
REDIS_CONF_NAME=redis-conf.yaml
# MySQL(gorm)配置文件名称
MYSQL_CONF_NAME=mysql-conf.yaml
# zerolog 配置文件名称
ZEROLOG_CONF_NAME=zerolog-conf.yaml
# ETCD 客户端配置文件名称 
ETCD_CLIENT_CONF_NAME=etcd-client3-conf.yaml
# SERVER 配置文件名称
SERVER_CONF_NAME=server-conf.yaml
# zap 配置文件名称
ZAP_CONF_NAME=zap-conf.yaml


echo 启动测试环境服务器...
go run main.go \
--redis-conf $CONFIG_ROOT_PATH/$CONFIG_APPLY_ENV/$REDIS_CONF_NAME \
--mysql-conf $CONFIG_ROOT_PATH/$CONFIG_APPLY_ENV/$MYSQL_CONF_NAME \
# --zerolog-conf $CONFIG_ROOT_PATH/$CONFIG_APPLY_ENV/$ZEROLOG_CONF_NAME \
# --etcd-client3-conf $CONFIG_ROOT_PATH/$CONFIG_APPLY_ENV/$ETCD_CLIENT_CONF_NAME \
# --server-conf $CONFIG_ROOT_PATH/$CONFIG_APPLY_ENV/$SERVER_CONF_NAME \
# --zap-conf $CONFIG_ROOT_PATH/$CONFIG_APPLY_ENV/$ZAP_CONF_NAME
