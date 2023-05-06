package main

import (
	"fmt"
	"os"
	"os/exec"
	"plugin"
	"time"
)

type Plugin struct{}

func (Plugin) DoSomething() string {
	return ""
}

func main() {
	p, err := plugin.Open("./plugins/plugin.so")
	if err != nil {
		fmt.Printf("failed to open plugin: %s\n", err)
		os.Exit(-1)
	}

	var s Plugin
	symPlugin, err := p.Lookup("Plugin")
	if err != nil {
		fmt.Printf("failed to lookup symbol: %s\n", err)
		os.Exit(-1)
	}

	s, ok := symPlugin.(Plugin)
	if !ok {
		fmt.Printf("invalid plugin symbol\n")
		os.Exit(-1)
	}

	t := time.NewTicker(10 * time.Second)

	for range t.C {
		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", "./plugins/plugin.so", "./plugins/plugin.go")
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
			continue
		}

		p, err = plugin.Open("./plugins/plugin.so")
		if err != nil {
			fmt.Printf("failed to open plugin: %s\n", err)
			continue
		}

		var news Plugin
		symPlugin, err := p.Lookup("Plugin")
		if err != nil {
			fmt.Printf("failed to lookup symbol: %s\n", err)
			continue
		}

		news, ok = symPlugin.(Plugin)
		if !ok {
			fmt.Printf("invalid plugin symbol\n")
			continue
		}

		s = news

		fmt.Println(s.DoSomething())
	}
}
