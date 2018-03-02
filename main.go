package main

import (
	"fmt"
	"os"

	"github.com/servntire/car-ownership/blockchain"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Channel parameters
		ChannelID:     "chainhero",
		ChannelConfig: "" + os.Getenv("GOPATH") + "/src/github.com/chainHero/heroes-service/fixtures/artifacts/",

		// Chaincode parameters
		OrgAdmin:   "Admin",
		OrgName:    "Org1",
		ConfigFile: "config.yaml",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
	}
}
