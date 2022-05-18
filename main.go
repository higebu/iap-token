package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apstndb/adcplus/tokensource"
)

var (
	iapClientID string
)

func init() {
	iapClientID = os.Getenv("IAP_CLIENT_ID")
}

func main() {
	ctx := context.Background()
	ts, err := tokensource.SmartIDTokenSource(ctx, iapClientID)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	token, err := ts.Token()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf(token.AccessToken)
}
