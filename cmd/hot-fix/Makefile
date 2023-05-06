SOURCES=main.go plugins/plugin.go
PLUGIN_SO=plugins/plugin.so
MAIN=main

.PHONY: all clean

all: $(MAIN)

$(MAIN): $(SOURCES)
	go build -buildmode=plugin -o $(PLUGIN_SO) plugins/plugin.go
	go build -o $(MAIN) .

clean:
	rm -rf $(PLUGIN_SO) $(MAIN)