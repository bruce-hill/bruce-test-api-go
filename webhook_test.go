// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/bruce-test-api-go"
	"github.com/stainless-sdks/bruce-test-api-go/internal/testutil"
	"github.com/stainless-sdks/bruce-test-api-go/option"
)

func TestWebhookRegister(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := brucetestapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Webhooks.Register(context.TODO(), brucetestapi.WebhookRegisterParams{
		URL: "https://example.com",
	})
	if err != nil {
		var apierr *brucetestapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
