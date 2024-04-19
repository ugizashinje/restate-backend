package main

import (
	"warrant-api/pkg/config"
	"warrant-api/pkg/db"
	"warrant-api/pkg/server"
	"warrant-api/pkg/superset"
	"warrant-api/pkg/wire"
)

func init() {
	config.Init("dev")
	wire.Svc = wire.Init("dev")

}
func main() {
	// testDB(gormDB)
	gormDB, err := db.Init(config.Db)
	if err != nil {
		panic(err)
	}

	c := make(chan int)
	go func() {
		superset.MaintainToken(c)
	}()
	<-c
	server.Start(gormDB)

}
