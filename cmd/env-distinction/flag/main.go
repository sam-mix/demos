package main

import (
	"flag"
	"fmt"
)

var env = flag.String("env", "dev", "set the environment")
var dbAddr, dbPort, logLevel string

func main() {
	flag.Parse()
	switch *env {
	case "dev":
		dbAddr = "localhost"
		dbPort = "3306"
		logLevel = "debug"
	case "test":
		dbAddr = "test.host.com"
		dbPort = "3306"
		logLevel = "info"
	case "pre":
		dbAddr = "pre.host.com"
		dbPort = "3307"
		logLevel = "warn"
	case "prod":
		dbAddr = "prod.host.com"
		dbPort = "3308"
		logLevel = "error"
	}
	// do something with the config, such as connecting database, setting logger level etc.
	fmt.Printf("current environment is %s, database address is %s:%s, log level is %s\n", *env, dbAddr, dbPort, logLevel)
}
