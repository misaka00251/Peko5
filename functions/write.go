package functions

import (
	"Peko5/config"
	"context"
	"fmt"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"log"
	"time"
)

var deleteEmailID string

// WriteEmailDraft is to write a email draft
func WriteEmailDraft(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	randString := GenerateRamdomString(16)
	requestBody := graphmodels.NewMessage()
	subject := "Executed by Peko5 on" + time.Now().Format("2006-01-02 15:04:05")
	requestBody.SetSubject(&subject)
	importance := graphmodels.LOW_IMPORTANCE
	requestBody.SetImportance(&importance)
	body := graphmodels.NewItemBody()
	contentType := graphmodels.HTML_BODYTYPE
	body.SetContentType(&contentType)
	content := `This is a message generated by Peko5. <br>
If you see this mail in your inbox, you should ignore and delete it. <br>
<b>Unique identifier: ` + randString + `</b><br>`
	body.SetContent(&content)
	requestBody.SetBody(body)

	recipient := graphmodels.NewRecipient()
	emailAddress := graphmodels.NewEmailAddress()
	// Send to yourself
	address := userEmail
	emailAddress.SetAddress(&address)
	recipient.SetEmailAddress(emailAddress)

	toRecipients := []graphmodels.Recipientable{
		recipient,
	}
	requestBody.SetToRecipients(toRecipients)

	messages, err := graphClient.Users().ByUserId(c.AdminID).Messages().Post(context.Background(), requestBody, nil)
	if err != nil {
		log.Fatalf("Error creating message: %v\n", err)
		return
	}
	fmt.Printf("Message ID: %v\n", *messages.GetId())
	deleteEmailID = *messages.GetId()
}

func DeleteEmailDraft(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	err := graphClient.Users().ByUserId(c.AdminID).Messages().ByMessageId(deleteEmailID).Delete(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error deleting message: %v\n", err)
		return
	}
	fmt.Println("Message deleted.")
}
