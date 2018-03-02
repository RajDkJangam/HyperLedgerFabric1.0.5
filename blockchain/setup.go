package blockchain

import (
	"fmt"
	"time"

	chmgmt "github.com/hyperledger/fabric-sdk-go/api/apitxn/chmgmtclient"
	"github.com/hyperledger/fabric-sdk-go/pkg/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// FabricSetup implementation
type FabricSetup struct {
	ConfigFile    string
	OrgID         string
	ChannelID     string
	initialized   bool
	ChannelConfig string
	OrgAdmin      string
	OrgName       string
}

// Initialize reads the configuration file and sets up the client, chain and event hub
func (setup *FabricSetup) Initialize() error {

	// Add parameters for the initialization
	if setup.initialized {
		return fmt.Errorf("sdk already initialized")
	}

	//TODO
	err := fmt.Errorf("")

	setup.sdk, err = fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return fmt.Errorf("failed to create sdk: %v", err)
	}

	// Channel management client is responsible for managing channels (create/update channel)
	// Supply user that has privileges to create channel (in this case orderer admin)
	chMgmtClient, err := setup.sdk.NewClient(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName)).ChannelMgmt()
	if err != nil {
		return fmt.Errorf("failed to add Admin user to sdk: %v", err)
	}

	// Org admin user is signing user for creating channel
	session, err := setup.sdk.NewClient(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName)).Session()
	if err != nil {
		return fmt.Errorf("failed to get session for %s, %s: %s", setup.OrgName, setup.OrgAdmin, err)
	}
	orgAdminUser := session

	// Create channel
	req := chmgmt.SaveChannelRequest{ChannelID: setup.ChannelID, ChannelConfig: setup.ChannelConfig + "chainhero.channel.tx", SigningIdentity: orgAdminUser}
	if err = chMgmtClient.SaveChannel(req); err != nil {
		return fmt.Errorf("failed to create channel: %v", err)
	}

	// Allow orderer to process channel creation
	time.Sleep(time.Second * 5)

	// Org resource management client
	setup.admin, err = setup.sdk.NewClient(fabsdk.WithUser(setup.OrgAdmin)).ResourceMgmt()
	if err != nil {
		return fmt.Errorf("failed to create new resource management client: %v", err)
	}

	// Org peers join channel
	if err = setup.admin.JoinChannel(setup.ChannelID); err != nil {
		return fmt.Errorf("org peers failed to join the channel: %v", err)
	}

	fmt.Println("Initialization Successful")
	return nil
}
