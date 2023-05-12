package setup

import (
	"flag"

	"github.com/spf13/viper"
)

func init() {
	// 热更
	configFile := flag.String("env-conf", "", "环境配置文件路径")
	// env := flag.String("env", "", "环境配置文件路径")
	flag.Parse()

	viper.SetConfigFile(*configFile)
	viper.ReadInConfig()
	// redisConfFile := viper.GetString(fmt.Sprintf("%s.redis-conf", *env))
	// mysqlConfFile := viper.GetString(fmt.Sprintf("%s.mysql-conf", *env))
	// zerologConfFile := viper.GetString(fmt.Sprintf("%s.zerolog-conf", *env))
	// etcdConfFile := viper.GetString(fmt.Sprintf("%s.etcd-client3-conf", *env))
	// serverConfFile := viper.GetString(fmt.Sprintf("%s.server-conf", *env))
	// zapConfFile := viper.GetString(fmt.Sprintf("%s.zap-conf", *env))

}
