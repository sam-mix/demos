// plugins/plugin.go

package main

import "time"

type Plugin struct{}

func (p Plugin) DoSomething() string {
	return "Hello, World! " + time.Now().String()
}
