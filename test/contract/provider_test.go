package contract

import (
	"github.com/pact-foundation/pact-go/dsl"
	"testing"
	"github.com/pact-foundation/pact-go/types"
	"path/filepath"
)

func TestProvider(t *testing.T) {
	pact := &dsl.Pact{
		Port:     6666,
	}

	path, _ := filepath.Abs("./apiGateway-helloWorld.json")

	// Verify the Provider with local Pact Files
	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:        "https://helloworld.dev.danielspeixoto.com",
		PactURLs:               []string{path},
	})
}