package main

import (
	"fmt"

	"github.com/spf13/viper"
)

var dbAddr, dbPort, logLevel string

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %s", err))
	}

	env := viper.GetString("environment")

	switch env {
	case "dev":
		dbAddr = viper.GetString("dev.db.addr")
		dbPort = viper.GetString("dev.db.port")
		logLevel = viper.GetString("dev.log.level")
	case "test":
		dbAddr = viper.GetString("test.db.addr")
		dbPort = viper.GetString("test.db.port")
		logLevel = viper.GetString("test.log.level")
	case "pre":
		dbAddr = viper.GetString("pre.db.addr")
		dbPort = viper.GetString("pre.db.port")
		logLevel = viper.GetString("pre.log.level")
	case "prod":
		dbAddr = viper.GetString("prod.db.addr")
		dbPort = viper.GetString("prod.db.port")
		logLevel = viper.GetString("prod.log.level")
	default:
		dbAddr = viper.GetString("dev.db.addr")
		dbPort = viper.GetString("dev.db.port")
		logLevel = viper.GetString("dev.log.level")
	}

	// do something with the config, such as connecting database, setting logger level etc.
	fmt.Printf("current environment is %s, database address is %s:%s, log level is %s\n", env, dbAddr, dbPort, logLevel)
}
