package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.reefassistant.com/apex"

	v1alpha1 "go.reefassistant.com/apex/client/private/v1alpha1"
)

const (
	APEX_URL      = "APEX_URL"
	APEX_USERNAME = "APEX_USERNAME"
	APEX_PASSWORD = "APEX_PASSWORD"
)

func main() {
	client, err := apex.NewDefaultClient(os.Getenv(APEX_URL))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Fetch status using public API.
	resp1, _, err := client.PublicV1().DefaultApi.GetStatus(ctx).Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("software: %q\nserial: %q\ninputs: %d\noutputs: %d\n",
		resp1.Istat.Software,
		resp1.Istat.Serial,
		len(resp1.Istat.Inputs),
		len(resp1.Istat.Outputs),
	)

	// Perform a login using the private API.
	resp2, _, err := client.PrivateV1().DefaultApi.LoginUser(ctx).LoginRequest(
		v1alpha1.LoginRequest{
			Login:    os.Getenv(APEX_USERNAME),
			Password: os.Getenv(APEX_PASSWORD),
		},
	).Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("session id: %q\n", resp2.ConnectSid)

	// Load user information using the private API.
	resp3, _, err := client.PrivateV1().DefaultApi.GetUser(ctx).Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("confirmed login as: %q\n", resp3.Username)
}
