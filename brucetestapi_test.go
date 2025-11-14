// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi_test

import (
	"bytes"
	"context"
	"errors"
	"io"
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
			FirstPos:    "first_pos",
			FirstQuery:  []int64{0},
			SecondQuery: brucetestapi.String("second_query"),
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
			FirstPos:   "first_pos",
			ArrayItems: []int64{0},
			Name: brucetestapi.PostFnordParamsName{
				FullName: "full_name",
				Nickname: brucetestapi.String("nickname"),
			},
			SecondQuery: brucetestapi.String("second_query"),
			ImageBase64: brucetestapi.String("U3RhaW5sZXNzIHJvY2tz"),
			ImageBinary: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			Job:         brucetestapi.String("job"),
			Pets: []brucetestapi.PostFnordParamsPet{{
				Name: brucetestapi.PostFnordParamsPetName{
					FullName: "full_name",
					Nickname: brucetestapi.String("nickname"),
				},
				Species: "species",
			}},
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

func TestBrucetestapiTestFormWithOptionalParams(t *testing.T) {
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
	_, err := client.TestForm(context.TODO(), brucetestapi.TestFormParams{
		Email:     "email",
		Username:  "U3RhaW5sZXNzIHJvY2tz",
		Age:       brucetestapi.Int(0),
		Subscribe: brucetestapi.Bool(true),
	})
	if err != nil {
		var apierr *brucetestapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
