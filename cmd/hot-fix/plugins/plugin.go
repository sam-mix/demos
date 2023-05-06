// plugins/plugin.go

package main

type Plugin struct{}

func (p Plugin) DoSomething() string {
	return "Hello, World!"
}
