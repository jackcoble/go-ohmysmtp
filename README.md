<div align="center">
	<h1>go-ohmysmtp</h1>
	A Go wrapper for the <a href="https://ohmysmtp.com">OhMySMTP</a> email service.
</div>

## Package
```
github.com/jackcoble/go-ohmysmtp
```

## Examples
Send an email.

```go
package main

import (
  "log"
  "github.com/jackcoble/go-ohmysmtp"
)

func main() {
	// Create a new API client using API key
	client := ohmysmtp.NewClient("MY_API_KEY")

	// Construct a payload
	payload := ohmysmtp.Payload{
		From: "hello@example.com",
		To: "test@example.com",
		Subject: "Sending an email using OhMySMTP!",
		HTMLBody: "<h1>Body that is HTML</h1>",
	}
	
	// Send the email
	if err := client.Send(payload); err != nil {
		log.Fatal(err)
	}
}
```