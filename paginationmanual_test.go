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

func TestManualPagination(t *testing.T) {
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
	page, err := client.Pagination.List(context.TODO(), brucetestapi.PaginationListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, pagination := range page.Items {
		t.Logf("%+v\n", pagination.Baz)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, pagination := range page.Items {
			t.Logf("%+v\n", pagination.Baz)
		}
	}
}
