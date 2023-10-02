package functions

import (
	"Peko5/config"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

// ReadActions contains all read actions
func ReadActions(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	GetAdmin(graphClient, c)
	GetDefaultGroup(graphClient, c)
	GetDefaultSite(graphClient, c)
	//BreakTime()
	GetAdminCalendar(graphClient, c)
	GetAdminOnedrive(graphClient, c)
	GetAdminEmail(graphClient, c)
	GetAdminOneNote(graphClient, c)
}

// WriteActions contains all write actions
func WriteActions(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	WriteEmailDraft(graphClient, c)
	//BreakTime()
	DeleteEmailDraft(graphClient, c)
}

// ActionList contains two parts: Read and Write
func ActionList(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	// Notice: Read can't get you renewed nowadays, you need to write stuff.
	ReadActions(graphClient, c)
	//BreakTime()
	WriteActions(graphClient, c)
}
