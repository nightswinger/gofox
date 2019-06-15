# gofox #

gofox is a Go client library for accessing the [Sigfox API(2.0)](https://support.sigfox.com/apidocs).

## Install ##

```bash
go get -u github.com/nightswinger/gofox/sigfox
```

## Usage ##

```go
import "github.com/nightswinger/gofox/sigfox"
```

Construct a new Sigfox client, then use services on the client to access different parts of the Sigfox API. For example:

```go
client := sigfox.NewClient("API_LOGIN_ID", "API_PASSWORD")

// Get device list
list, err := client.Device.List(nil)

// Get device messages
msg, err := client.Device.Messages("DeviceID")

// Get device type information with context
ctx := context.Background()
info, err := client.DeviceType.InfoContext(ctx, "DeviceID")
```
