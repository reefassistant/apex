package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.reefassistant.com/apex"
)

func main() {
	client, err := apex.NewDefaultClient(os.Getenv("APEX_HOSTNAME"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, _, err := client.PublicV1().DefaultApi.GetStatus(ctx).Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("software: %q\nserial: %q\ninputs: %d\noutputs: %d\n",
		resp.Istat.Software,
		resp.Istat.Serial,
		len(resp.Istat.Inputs),
		len(resp.Istat.Outputs),
	)
}
