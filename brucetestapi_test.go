// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/bruce-hill/bruce-test-api-go"
	"github.com/bruce-hill/bruce-test-api-go/internal/testutil"
	"github.com/bruce-hill/bruce-test-api-go/option"
)

func TestBrucetestapiFormTestWithOptionalParams(t *testing.T) {
	t.Skip("prism issues because it's bad")
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
	_, err := client.FormTest(
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
			IDOrIndex: brucetestapi.FormTestParamsIDOrIndexUnion{
				OfInt: brucetestapi.Int(0),
			},
			Limit: brucetestapi.Int(1),
			Tags:  []string{"string"},
			Blorp: brucetestapi.String("example value"),
			ManySomethings: []brucetestapi.FormTestParamsManySomethingUnion{{
				OfFloat: brucetestapi.Float(123),
			}, {
				OfThingy: &brucetestapi.FormTestParamsManySomethingThingy{
					Name:  "Bozo",
					Count: brucetestapi.Int(5),
				},
			}, {
				OfThingy: &brucetestapi.FormTestParamsManySomethingThingy{
					Name:  "Fran",
					Count: brucetestapi.Int(5),
				},
			}},
			Pets: []brucetestapi.FormTestParamsPet{{
				Name: "Alfie",
				Age:  brucetestapi.Int(0),
			}, {
				Name: "Brando",
				Age:  brucetestapi.Int(12),
			}, {
				Name: "Charlie",
				Age:  brucetestapi.Int(0),
			}},
			PlsNull: nil,
			Preferences: brucetestapi.FormTestParamsPreferences{
				Alerts: brucetestapi.Bool(true),
				Theme:  brucetestapi.String("dark"),
			},
			Something: brucetestapi.FormTestParamsSomethingUnion{
				OfThingy: &brucetestapi.FormTestParamsSomethingThingy{
					Name:  "Albert",
					Count: brucetestapi.Int(5),
				},
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
	_, err := client.JsonTest(
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
			IDOrIndex: brucetestapi.JsonTestParamsIDOrIndexUnion{
				OfInt: brucetestapi.Int(0),
			},
			Limit: brucetestapi.Int(1),
			Tags:  []string{"string"},
			Blorp: brucetestapi.String("test data"),
			ManySomethings: []brucetestapi.JsonTestParamsManySomethingUnion{{
				OfFloat: brucetestapi.Float(123),
			}, {
				OfThingy: &brucetestapi.JsonTestParamsManySomethingThingy{
					Name:  "Bozo",
					Count: brucetestapi.Int(5),
				},
			}, {
				OfThingy: &brucetestapi.JsonTestParamsManySomethingThingy{
					Name:  "Fran",
					Count: brucetestapi.Int(5),
				},
			}},
			Pets: []brucetestapi.JsonTestParamsPet{{
				Name: "Alfie",
				Age:  brucetestapi.Int(0),
			}, {
				Name: "Brando",
				Age:  brucetestapi.Int(12),
			}, {
				Name: "Charlie",
				Age:  brucetestapi.Int(0),
			}},
			PlsNull: nil,
			Preferences: brucetestapi.JsonTestParamsPreferences{
				Alerts: brucetestapi.Bool(false),
				Theme:  brucetestapi.String("light"),
			},
			Something: brucetestapi.JsonTestParamsSomethingUnion{
				OfThingy: &brucetestapi.JsonTestParamsSomethingThingy{
					Name:  "Albert",
					Count: brucetestapi.Int(5),
				},
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
