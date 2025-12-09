// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/DefinitelyATestOrg/test-api-go"
	"github.com/DefinitelyATestOrg/test-api-go/internal/testutil"
	"github.com/DefinitelyATestOrg/test-api-go/option"
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
	err := client.FormTest(
		context.TODO(),
		"abc123",
		brucetestapi.FormTestParams{
			Version:  1,
			Date:     time.Now(),
			Datetime: time.Now(),
			Time:     "18:11:19.117Z",
			Filter: brucetestapi.FormTestParamsFilter{
				Status: brucetestapi.String("active"),
				Meta: brucetestapi.FormTestParamsFilterMeta{
					Level: brucetestapi.Int(3),
				},
			},
			Limit: brucetestapi.Int(20),
			Tags:  []string{"red", "large"},
			Preferences: brucetestapi.FormTestParamsPreferences{
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
