// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/bruce-test-api-go/internal/apijson"
	"github.com/stainless-sdks/bruce-test-api-go/internal/requestconfig"
	"github.com/stainless-sdks/bruce-test-api-go/packages/param"
	"github.com/stainless-sdks/bruce-test-api-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PageNumber[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
	Pages int64 `json:"pages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Total       respjson.Field
		Page        respjson.Field
		Size        respjson.Field
		Pages       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PageNumber[T]) RawJSON() string { return r.JSON.raw }
func (r *PageNumber[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageNumber[T]) GetNextPage() (res *PageNumber[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	currentPage := r.Page
	if currentPage >= r.Pages {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PageNumber[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageNumber[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageNumberAutoPager[T any] struct {
	page *PageNumber[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPageNumberAutoPager[T any](page *PageNumber[T], err error) *PageNumberAutoPager[T] {
	return &PageNumberAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageNumberAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageNumberAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageNumberAutoPager[T]) Err() error {
	return r.err
}

func (r *PageNumberAutoPager[T]) Index() int {
	return r.run
}
