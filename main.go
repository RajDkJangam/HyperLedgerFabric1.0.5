package main

import (
	"fmt"
	"os"
	"github.com/servntire/car-ownership/blockchain"
	"github.com/servntire/car-ownership/web/controllers"
	"github.com/servntire/car-ownership/web"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Channel parameters
		ChannelID:        	"chainhero",
		ChannelConfig:    	"" + os.Getenv("GOPATH") + "/src/github.com/servntire/car-ownership/fixtures/artifacts/",

		// Chaincode parameters
		ChainCodeID:      	"heroes-service",
		ChaincodeGoPath:  	os.Getenv("GOPATH"),
		ChaincodePath:    	"github.com/servntire/car-ownership/chaincode/",
		OrgAdmin:			"Admin",
		OrgName:			"Org1",
		ConfigFile:			"config.yaml",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
	}

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
	}

	// Launch the web application listening
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)

}
