package main

// func main() {
// 	s := server.NewServer()
// 	r := &serverplugin.EtcdV3RegisterPlugin{
// 		ServiceAddress:   "tcp@localhost:8972",
// 		RegisterInterval: time.Minute,
// 		TTL:              time.Minute * 3,
// 		BasePath:         "/services/",
// 		EtcdServers:      []string{"http://localhost:2379"},
// 	}
// 	err := r.Start()
// 	if err != nil {
// 		log.Fatalf("failed to start etcd registry plugin: %v", err)
// 	}
// 	s.Plugins.Add(r)
// 	s.RegisterName("HelloService", new(HelloService), "")
// 	s.Serve("tcp", ":8972")
// }
