package main

import (
	"echo-api-example/blockchain"
	"echo-api-example/server"
	"echo-api-example/utils"
)

func main() {
	utils.InitEnvConfigs()
	
	client, err := blockchain.NewEthClient(utils.EnvConfigs.ETH_rpc)
	if err != nil {
		panic("Eth client cannot be created")
	}


	server.Start(client)
}
