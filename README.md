This POC has 2 components (Server and Agent Service).
Server is an http server which serves /policies on port 3040 (http://localhost:3040/policies).
Agent Service is a windows service, which on startup fetches policies and applies on windows machine.

## Server
1. Start the server - `go run cmd/server/main.go`
or
2. Build and run .exe
    a. `go build -o dist/mdm-server.exe ./cmd/server`
    b. `./dist/mdm-server.exe`

## Agent Service

### Build the agent service
1. From windows - `go build -o dist/mdm-agent-service.exe ./cmd/agent`
2. From different OS - `GOOS=windows GOARCH=amd64 go build -o dist/mdm-agent-service.exe ./cmd/agent`

### Install/Uninstall the service
1. Open Command Prompt as Admin
2. Create service - `sc create "MDM Agent Service POC" binPath="{path of above binary bulild}" start=auto`
    eg. `sc create "MDM Agent Service POC" binPath="C:\go\src\github.com\kdkumawat\mdm-poc\dist\mdm-agent-service.exe" start=auto`
3. Start service - `sc start "MDM Agent Service POC"`
4. Stop service - `sc stop "MDM Agent Service POC"`
5. Delete service - `sc delete "MDM Agent Service POC"`


## Verify Agent Service is applying policies

### Option 1 (view in Registry Editor)
1. Start + Run
2. type `regedit`
3. clik Ok
4. Navigate to `Computer\HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`
5. Verify Data for `LegalNoticeCaption` and `LegalNoticeText`

### Option 2 (Restart / Sign out from current user)
1. On login screen
2. User should see applied Caption and Text
