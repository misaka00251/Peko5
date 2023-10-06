package functions

import (
	"Peko5/config"
	"context"
	"fmt"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"log"
)

var userEmail string

// GetAdmin is to get the admin's information
func GetAdmin(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	users, err := graphClient.Users().ByUserId(c.AdminID).Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting user: %v\n", err)
		return
	}
	fmt.Printf("Login with: %v <%v>\n", *users.GetDisplayName(), *users.GetMail())
	userEmail = *users.GetMail()
}

// GetAdminOnedrive is to get the admin's onedrive information
func GetAdminOnedrive(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	drive, err := graphClient.Users().ByUserId(c.AdminID).Drive().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting drive: %v\n", err)
		return
	}
	fmt.Printf("Drive Size: %.2f GB\n", float64(*drive.GetQuota().GetTotal())/float64(1000000000))
	fmt.Printf("Drive ID: %v\n", *drive.GetId())

	recent, err := graphClient.Drives().ByDriveId(*drive.GetId()).Recent().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting drive recent files: %v\n", err)
		return
	}
	fmt.Printf("Recent file ID: %v\n", *recent.GetValue()[0].GetId())
}

// GetDefaultGroup is to get the default group's information
func GetDefaultGroup(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	groups, err := graphClient.Groups().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting groups: %v\n", err)
		return
	}
	fmt.Printf("Default group ID: %v\n", *groups.GetValue()[0].GetId())
}

// GetDefaultSite is to get the default site's information
func GetDefaultSite(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	lists, err := graphClient.Sites().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting groups: %v\n", err)
		return
	}
	fmt.Printf("Default SharePoint site ID: %v\n", *lists.GetValue()[0].GetId())

	drives, err := graphClient.Sites().BySiteId(*lists.GetValue()[0].GetId()).Drives().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting groups: %v\n", err)
		return
	}
	fmt.Printf("Default SharePoint site drive ID: %v\n", *drives.GetValue()[0].GetId())

	columns, err := graphClient.Sites().BySiteId(*lists.GetValue()[0].GetId()).Columns().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting groups: %v\n", err)
		return
	}
	fmt.Printf("Core Contact and Calendar Columns ID: %v\n", *columns.GetValue()[0].GetId())
}

// GetAdminCalendar is to get the admin's calendar information
func GetAdminCalendar(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	calendars, err := graphClient.Users().ByUserId(c.AdminID).Calendars().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting calendars: %v\n", err)
		return
	}
	fmt.Printf("Default calendar ID: %v\n", *calendars.GetValue()[0].GetId())
}

// GetAdminEmail is to get the admin's email information
func GetAdminEmail(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	masterCategories, err := graphClient.Users().ByUserId(c.AdminID).Outlook().MasterCategories().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting emails: %v\n", err)
		return
	}
	fmt.Printf("Red category ID: %v\n", *masterCategories.GetValue()[0].GetId())
}

// GetAdminOneNote is to get the admin's OneNote information
func GetAdminOneNote(graphClient *msgraphsdk.GraphServiceClient, c config.Config) {
	notebooks, err := graphClient.Users().ByUserId(c.AdminID).Onenote().Sections().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting OneNote: %v\n", err)
		return
	}
	fmt.Printf("Default OneNote section ID: %v\n", *notebooks.GetValue()[0].GetId())
}
