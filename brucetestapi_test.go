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

func TestBrucetestapiFnordWithOptionalParams(t *testing.T) {
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
	_, err := client.Fnord(
		context.TODO(),
		"second_pos",
		brucetestapi.FnordParams{
			FirstPos:   "first_pos",
			FirstQuery: []int64{0},
			SecondQuery: brucetestapi.FnordParamsSecondQuery{
				Age:  brucetestapi.Int(0),
				Name: brucetestapi.String("name"),
			},
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

func TestBrucetestapiPostFnordWithOptionalParams(t *testing.T) {
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
	_, err := client.PostFnord(
		context.TODO(),
		"second_pos",
		brucetestapi.PostFnordParams{
			FirstPos:    "first_pos",
			FirstQuery:  []int64{0},
			FullName:    "full_name",
			SecondQuery: brucetestapi.String("second_query"),
			Nickname:    brucetestapi.String("nickname"),
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
