package main

import (
	"fmt"

	"github.com/op/go-logging"
	"brilliance/client_e2e_test/blockchain/backend/service"
	"brilliance/client_e2e_test/blockchain/backend/service/fabric"
	"brilliance/client_e2e_test/blockchain/common/config"
	"brilliance/client_e2e_test/blockchain/database/mongo"
)

var (
	configPath = "config/config.yaml"
)

func initFab() {
	config.InitConfig(configPath)
}

var Log = logging.MustGetLogger("turbot")

func main() {
	mongo.OpenMongo("10.0.90.152:21117")
	initFab()
	var err error
	service.FabClient, err = fabric.NewFabricClient()
	if err != nil {
		fmt.Println(fmt.Errorf("error to get client", err))
		return
	}
	initRouter().Run(":8889")
}
