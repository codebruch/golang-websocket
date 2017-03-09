# golang-websocket
About

Simple app to test web socket connections passed into the application.

Instructions to deploy as this application.

Run the following commands

```
git clone https://github.com/cf-routing/golang-websocket.git
cd golang-websocket
cf api api.domain.com
cf auth [your_username] [your_password]
cf create-org testorg
cf target -o testorg
cf create-space testspace
cf target -o testorg -s testspace
cf push golang
```

After staging the app successfully, run the command `cf app golang` to get the URL for the app (e.g. golang.domain.com).
- `cf logs golang` 
- Open your app in the browser with the `index` path (e.g. `golang.domain.com/index`)
- Hit the "Send Message" button in the browser
- If websockets are configured correctly in your CF, there should be logs with the message you sent
