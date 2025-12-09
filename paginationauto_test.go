// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/DefinitelyATestOrg/test-api-go"
	"github.com/DefinitelyATestOrg/test-api-go/internal/testutil"
	"github.com/DefinitelyATestOrg/test-api-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.ListFoosAutoPaging(context.TODO(), brucetestapi.ListFoosParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		client := iter.Current()
		t.Logf("%+v\n", client.Baz)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
