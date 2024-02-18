package main

import (
	"echo-api-example/blockchain"
	"echo-api-example/server"
	"echo-api-example/utils"
	"fmt"
)

func main() {
	utils.InitEnvConfigs()
	fmt.Printf(">>>>>> %v", utils.EnvConfigs.ETH_rpc)
	
	client, err := blockchain.NewEthClient("TODO_ADD_RPC_URL")
	if err != nil {
		panic("Eth client cannot be created")
	}


	server.Start(client)
}
