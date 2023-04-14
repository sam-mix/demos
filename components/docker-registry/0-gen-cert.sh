#!/usr/bin/env bash


# 创建 CA 私钥
openssl genrsa -out "root-ca.key" 4096

# 利用私钥创建 CA 根证书请求文件
openssl req \
          -new -key "root-ca.key" \
          -out "root-ca.csr" -sha256 \
          -subj '/C=CN/ST=Guangdong/L=Shenzhen/O=yz/CN=yz Docker Registry CA'

# 配置 CA 根证书，新建 root-ca.cnf
tee root-ca.cnf <<EOF
[root_ca]
basicConstraints = critical,CA:TRUE,pathlen:1
keyUsage = critical, nonRepudiation, cRLSign, keyCertSign
subjectKeyIdentifier=hash
EOF

# 签发根证书
openssl x509 -req  -days 3650  -in "root-ca.csr" \
               -signkey "root-ca.key" -sha256 -out "root-ca.crt" \
               -extfile "root-ca.cnf" -extensions \
               root_ca

# 生成站点 SSL 私钥
openssl genrsa -out "docker.domain.com.key" 4096

# 使用私钥生成证书请求文件
openssl req -new -key "docker.domain.com.key" -out "site.csr" -sha256 \
          -subj '/C=CN/ST=Guangdong/L=Shenzhen/O=yz/CN=docker.domain.com'

# 配置证书，新建 site.cnf 文件
tee site.cnf <<EOF
[server]
authorityKeyIdentifier=keyid,issuer
basicConstraints = critical,CA:FALSE
extendedKeyUsage=serverAuth
keyUsage = critical, digitalSignature, keyEncipherment
subjectAltName = DNS:docker.domain.com, IP:127.0.0.1
subjectKeyIdentifier=hash
EOF

# 签署站点 SSL 证书
openssl x509 -req -days 750 -in "site.csr" -sha256 \
    -C" -CAkey "root-ca.key"  -CAcreateserial \
    -out "docker.domain.com.crt" -extfile "site.cnf" -extensions server

# 配置私有仓库
tee config.yml <<EOF
version: 0.1
log:
  accesslog:
    disabled: true
  level: debug
  formatter: text
  fields:
    service: registry
    environment: staging
storage:
  delete:
    enabled: true
  cache:
    blobdescriptor: inmemory
  filesystem:
    rootdirectory: /var/lib/registry
auth:
  htpasswd:
    realm: basic-realm
    path: /etc/docker/registry/auth/nginx.htpasswd
http:
  addr: :443
  host: https://docker.domain.com
  headers:
    X-Content-Type-Options: [nosniff]
  http2:
    disabled: false
  tls:
    certificate: /etc/docker/registry/ssl/docker.domain.com.crt
    key: /etc/docker/registry/ssl/docker.domain.com.key
health:
  storagedriver:
    enabled: true
    interval: 10s
threshold: 3

EOF


# 生成 http 认证文件
tee test.sh <<EOF
mkdir -p auth

docker run --rm \
    --entrypoint htpasswd \
    httpd:alpine \
    -Bbn username password > auth/nginx.htpasswd

EOF

chmod +x *.sh

# 添加假域名
echo "127.0.0.1 docker.domain.com" >> /etc/hosts

# 生成 docker-compose 文件
tee docker-compose.yml <<EOF
version: '3'

services:
  registry:
    image: registry
    ports:
      - "443:443"
    volumes:
      - ./:/etc/docker/registry
      - registry-data:/var/lib/registry

volumes:
  registry-data:

EOF



mkdir -p /etc/docker/certs.d/docker.domain.com
cp root-ca.crt /etc/docker/certs.d/docker.domain.com/ca.crt
