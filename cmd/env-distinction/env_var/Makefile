.PHONY: clean build test-dev test-test test-pre test-prod

clean:
	go clean
	rm -f config.yaml

build: clean
	env GOOS=linux GOARCH=amd64 go build -v main.go

test-dev:
	echo "environment: dev" > config.yaml
	echo "dev:" >> config.yaml
	echo "  db:" >> config.yaml
	echo "    addr: localhost" >> config.yaml
	echo "    port: 3306" >> config.yaml
	echo "  log:" >> config.yaml
	echo "    level: debug" >> config.yaml
	go run main.go

test-test:
	echo "environment: test" > config.yaml
	echo "test:" >> config.yaml
	echo "  db:" >> config.yaml
	echo "    addr: test.host.com" >> config.yaml
	echo "    port: 3306" >> config.yaml
	echo "  log:" >> config.yaml
	echo "    level: info" >> config.yaml
	go run main.go

test-pre:
	echo "environment: pre" > config.yaml
	echo "pre:" >> config.yaml
	echo "  db:" >> config.yaml
	echo "    addr: pre.host.com" >> config.yaml
	echo "    port: 3307" >> config.yaml
	echo "  log:" >> config.yaml
	echo "    level: warn" >> config.yaml
	go run main.go

test-prod:
	echo "environment: prod" > config.yaml
	echo "prod:" >> config.yaml
	echo "  db:" >> config.yaml
	echo "    addr: prod.host.com" >> config.yaml
	echo "    port: 3308" >> config.yaml
	echo "  log:" >> config.yaml
	echo "    level: error" >> config.yaml
	go run main.go