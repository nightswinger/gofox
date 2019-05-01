# gofox #

gofox is a Go client library for accessing the [Sigfox API(2.0)](https://support.sigfox.com/apidocs).

## Usage ##

```go
import "github.com/nightswinger/gofox/sigfox"
```

Construct a new Sigfox client, then use services on the client to access different parts of the Sigfox API. For example:

```go
client := sigfox.NewClient("API_LOGIN_ID", "API_PASSWORD")

// Get device messages
msg, _ := client.Device.Messages("DeviceID")
```