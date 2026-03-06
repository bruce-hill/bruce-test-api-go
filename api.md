# brucetestapi

Response Types:

- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#FormTestResponse">FormTestResponse</a>
- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#JsonTestResponse">JsonTestResponse</a>
- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#UpdateCountResponse">UpdateCountResponse</a>
- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#UploadTestResponse">UploadTestResponse</a>
- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#VersionResponse">VersionResponse</a>

Methods:

- <code title="delete /delete">client.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#BrucetestapiService.DeleteTest">DeleteTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /download">client.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#BrucetestapiService.DownloadTest">DownloadTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /form-v{version}/users/{userId}">client.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#BrucetestapiService.FormTest">FormTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#FormTestParams">FormTestParams</a>) (\*<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#FormTestResponse">FormTestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /json-v{version}/users/{userId}">client.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#BrucetestapiService.JsonTest">JsonTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#JsonTestParams">JsonTestParams</a>) (\*<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#JsonTestResponse">JsonTestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /nullable">client.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#BrucetestapiService.NullableTest">NullableTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#NullableTestParams">NullableTestParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /count">client.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#BrucetestapiService.UpdateCount">UpdateCount</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#UpdateCountParams">UpdateCountParams</a>) (\*<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#UpdateCountResponse">UpdateCountResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /upload">client.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#BrucetestapiService.UploadTest">UploadTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#UploadTestParams">UploadTestParams</a>) (\*<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#UploadTestResponse">UploadTestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /version">client.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#BrucetestapiService.Version">Version</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#VersionResponse">VersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /void-response">client.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#BrucetestapiService.VoidTest">VoidTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Pagination

Response Types:

- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#PaginationListResponse">PaginationListResponse</a>

Methods:

- <code title="get /paginated">client.Pagination.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#PaginationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#PaginationListParams">PaginationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2/packages/pagination#PageNumber">PageNumber</a>[<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#PaginationListResponse">PaginationListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ints

Methods:

- <code title="get /paginated-int">client.Pagination.Ints.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#PaginationIntService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#PaginationIntListParams">PaginationIntListParams</a>) (\*<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2/packages/pagination#PageNumber">PageNumber</a>[<a href="https://pkg.go.dev/builtin#int64">int64</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# StreamJson

Response Types:

- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#StreamJsonStreamResponse">StreamJsonStreamResponse</a>

Methods:

- <code title="get /stream-json">client.StreamJson.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#StreamJsonService.Stream">Stream</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/v2#StreamJsonStreamResponse">StreamJsonStreamResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
