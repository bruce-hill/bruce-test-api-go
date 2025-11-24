// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/bruce-test-api-go"
	"github.com/stainless-sdks/bruce-test-api-go/internal/testutil"
	"github.com/stainless-sdks/bruce-test-api-go/option"
)

func TestUsage(t *testing.T) {
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
	err := client.Foo(
		context.TODO(),
		"abc123",
		brucetestapi.FooParams{
			Version: 1,
			Filter: brucetestapi.FooParamsFilter{
				Status: brucetestapi.String("active"),
				Meta: brucetestapi.FooParamsFilterMeta{
					Level: brucetestapi.Int(3),
				},
			},
			Limit: brucetestapi.Int(20),
			Tags:  []string{"red", "large"},
			Preferences: brucetestapi.FooParamsPreferences{
				Theme:  brucetestapi.String("dark"),
				Alerts: brucetestapi.Bool(true),
			},
			XFlags:   []string{"fast", "debug", "verbose"},
			XTraceID: brucetestapi.String("trace-9f82b1"),
		},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
