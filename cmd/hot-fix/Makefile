.PHONY: all plugin main
all: plugin main

plugin:
    go build -buildmode=plugin -o plugins/plugin.so plugins/plugin.go

main:
     go build -o main main.go