.PHONY: gen-cert
gen-cert:
	@sh ./0-gen-cert.sh

.PHONY: tls-server
tls-server:
	@go run ./cmd/tcp_tls/server/main.go

.PHONY: tls-client
tls-client:
	@go run ./cmd/tcp_tls/client/main.go

.PHONY: geoip
geoip:
	@go run ./cmd/geoip/main.go

