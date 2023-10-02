# Peko5

This is a program to trigger your subscription API for you.

## Usage

0. Please use these services with your admin account at least once:

 - OneNote
 - OneDrive for Business (And upload something to it)

1. Sign in to the [Microsoft Entra admin center](https://entra.microsoft.com/), then click **Applications** -> **App registrations** on the left.
2. Click **New registration**.
3. Enter a name for your application, then click **Register**. Don't change anything else.
4. Goto your new app, click **API permissions** on the left, then click **Add a permission**.
5. Click **Microsoft Graph**. You need these permissions:

 - Calendars.Read
 - Files.Read.All
 - Group.Read.All
 - Mail.Read
 - Mail.ReadBasic
 - Mail.ReadBasic.All
 - Mail.ReadWrite
 - Mail.Send
 - Notes.Read.All
 - Sites.Read.All
 - User.Read.All

6. Click **Add permissions**.
7. Click **Grant admin consent for [your tenant]**, then click **Yes**.
8. Click **Certificates & secrets** on the left, then click **New client secret**.
9. Enter a description for your secret, then click **Add**.
10. Copy the value of the secret, write it down.
11. Click **Overview** on the left, then copy the value of **Application (client) ID**. Write it down.
12. Also copy the value of **Directory (tenant) ID**. Write it down.
13. Goto [Microsoft Entra admin center](https://entra.microsoft.com/), click **Users** -> **All users** on the left, then click on your admin account. Copy the value of **Object ID** and write it down.
14. Create **config.json** from template. Copy these into the config file.
15. Put your config.json in the same directory as peko5. Then run it.

### Run as a systemd service

We already provided a simple systemd service file.

## Disclaimer

This program is provided as-is, without any warranty. Use it at your own risk.

## License

WTFPL
