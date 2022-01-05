package job

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestService_GetJobs(t *testing.T) {
	pact := dsl.Pact{
		Provider:                 "go-provider",
		LogDir:                   "../../logs",
		PactDir:                  "../../pacts",
		DisableToolValidityCheck: true,
		LogLevel:                 "INFO",
	}

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", 8080),
		Tags:                       []string{"master"},
		FailIfNoPactsFound:         false,
		BrokerURL:                  "https://selahattinceylan.pactflow.io",
		BrokerToken:                "MGpB-0-JCU5yM8i3eUdA0Q",
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		StateHandlers:              stateHandlers,
		//RequestFilter:              fixBearerToken,
	})

	if err != nil {
		t.Fatal(err)
	}
}

var stateHandlers = types.StateHandlers{
	"a request to get jobs": func() error {

		return nil
	},
}
