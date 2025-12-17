// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/DefinitelyATestOrg/test-api-go"
	"github.com/DefinitelyATestOrg/test-api-go/internal/testutil"
	"github.com/DefinitelyATestOrg/test-api-go/option"
)

func TestBrucetestapiFormTestWithOptionalParams(t *testing.T) {
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
		"usr_abc123",
		brucetestapi.FormTestParams{
			Version:  2,
			Date:     time.Now(),
			Datetime: time.Now(),
			Time:     "18:11:19.117Z",
			Filter: brucetestapi.FormTestParamsFilter{
				Meta: brucetestapi.FormTestParamsFilterMeta{
					Level: brucetestapi.Int(0),
				},
				Status: brucetestapi.String("status"),
			},
			Limit: brucetestapi.Int(1),
			Tags:  []string{"string"},
			Blorp: brucetestapi.String("example value"),
			Preferences: brucetestapi.FormTestParamsPreferences{
				Alerts: brucetestapi.Bool(true),
				Theme:  brucetestapi.String("dark"),
			},
			XFlags:   []string{"feature-a", "feature-b"},
			XTraceID: brucetestapi.String("trace-xyz-789"),
		},
	)
	if err != nil {
		var apierr *brucetestapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrucetestapiJsonTestWithOptionalParams(t *testing.T) {
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
	err := client.JsonTest(
		context.TODO(),
		"usr_def456",
		brucetestapi.JsonTestParams{
			Version:  3,
			Date:     time.Now(),
			Datetime: time.Now(),
			Time:     "18:11:19.117Z",
			Filter: brucetestapi.JsonTestParamsFilter{
				Meta: brucetestapi.JsonTestParamsFilterMeta{
					Level: brucetestapi.Int(0),
				},
				Status: brucetestapi.String("status"),
			},
			Limit: brucetestapi.Int(1),
			Tags:  []string{"string"},
			Blorp: brucetestapi.String("test data"),
			Preferences: brucetestapi.JsonTestParamsPreferences{
				Alerts: brucetestapi.Bool(false),
				Theme:  brucetestapi.String("light"),
			},
			XFlags:   []string{"experimental", "verbose"},
			XTraceID: brucetestapi.String("trace-abc-123"),
		},
	)
	if err != nil {
		var apierr *brucetestapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrucetestapiUpdateCount(t *testing.T) {
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
	err := client.UpdateCount(context.TODO(), brucetestapi.UpdateCountParams{
		Body: 42,
	})
	if err != nil {
		var apierr *brucetestapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
