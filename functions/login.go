package functions

import (
	"Peko5/config"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

func InitProject(config config.Config) *msgraphsdk.GraphServiceClient {
	cred, err := azidentity.NewClientSecretCredential(
		config.TenantID,
		config.ClientID,
		config.ClientSecret,
		nil,
	)

	if err != nil {
		log.Fatalf("Error while login: %v\n", err)
	}

	graphClient, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"https://graph.microsoft.com/.default"})
	if err != nil {
		log.Fatalf("Error creating client: %v\n", err)
		return nil
	}

	return graphClient
}
