// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetestapi

import (
	"github.com/stainless-sdks/bruce-test-api-go/option"
)

// JsonTestService contains methods and other services that help with interacting
// with the bruce-test-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJsonTestService] method instead.
type JsonTestService struct {
	Options []option.RequestOption
}

// NewJsonTestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJsonTestService(opts ...option.RequestOption) (r JsonTestService) {
	r = JsonTestService{}
	r.Options = opts
	return
}
