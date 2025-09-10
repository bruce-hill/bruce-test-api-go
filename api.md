# brucetestapi

Response Types:

- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#JsonTestResponse">JsonTestResponse</a>

Methods:

- <code title="get /json-test">client.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#BrucetestapiService.JsonTest">JsonTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#JsonTestResponse">JsonTestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Text

Response Types:

- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#TextSetResponse">TextSetResponse</a>

Methods:

- <code title="put /foo-text">client.Text.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#TextService.Set">Set</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#TextSetParams">TextSetParams</a>) (<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#TextSetResponse">TextSetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Foos

Response Types:

- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#FooGetResponse">FooGetResponse</a>
- <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#FooListResponse">FooListResponse</a>

Methods:

- <code title="get /foo">client.Foos.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#FooService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#FooGetResponse">FooGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /all-foos">client.Foos.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#FooService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#FooListParams">FooListParams</a>) (<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go/packages/pagination#PageNumber">PageNumber</a>[<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/bruce-hill/bruce-test-api-go#FooListResponse">FooListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
