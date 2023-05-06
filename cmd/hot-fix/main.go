package main

import (
	"fmt"
	"os"
	"os/exec"
	"plugin"
	"time"
)

func main() {
	p, err := plugin.Open("./plugins/plugin.so")
	if err != nil {
		fmt.Printf("Failed to open plugin: %s\n", err)
		os.Exit(-1)
	}

	var s Plugin
	sSym, err := p.Lookup("Plugin")
	if err != nil {
		fmt.Printf("Failed to lookup plugin: %s\n", err)
		os.Exit(-1)
	}

	s, ok := sSym.(Plugin)
	if !ok {
		fmt.Printf("Invalid plugin symbol\n")
		os.Exit(-1)
	}

	fmt.Println(s.DoSomething())

	// 自动更新
	for {
		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", "./plugins/plugin.so", "./plugins/plugin.go")
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}

		p, err = plugin.Open("./plugins/plugin.so")
		if err != nil {
			fmt.Printf("Failed to open plugin: %s\n", err)
			continue
		}

		sSym, err = p.Lookup("Plugin")
		if err != nil {
			fmt.Printf("Failed to lookup plugin: %s\n", err)
			continue
		}

		s, ok = sSym.(Plugin)
		if !ok {
			fmt.Printf("Invalid plugin symbol\n")
			os.Exit(-1)
		}

		fmt.Println(s.DoSomething())

		time.Sleep(time.Second * 5)
	}
}

type Plugin interface {
	DoSomething() string
}
