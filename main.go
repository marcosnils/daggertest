package main

import (
	"context"
	"fmt"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
)

// adding something here

func main() {

	ctx := context.Background()

	keySet := oidc.NewRemoteKeySet(ctx, "https://token.actions.githubusercontent.com/.well-known/jwks")
	verifier := oidc.NewVerifier("https://token.actions.githubusercontent.com", keySet, &oidc.Config{SkipClientIDCheck: true})

	t, err := verifier.Verify(ctx, os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN"))

	fmt.Printf("Result is: %+v - %v\n", t, err)

}
