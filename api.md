# brucetestapi

Methods:

- <code title="post /form-v{version}/users/{userId}">client.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go#BrucetestapiService.FormTest">FormTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go#FormTestParams">FormTestParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /json-v{version}/users/{userId}">client.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go#BrucetestapiService.JsonTest">JsonTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go#JsonTestParams">JsonTestParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /count">client.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go#BrucetestapiService.UpdateCount">UpdateCount</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go#UpdateCountParams">UpdateCountParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Foos

Response Types:

- <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go#FooListResponse">FooListResponse</a>

Methods:

- <code title="get /paginated">client.Foos.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go#FooService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go#FooListParams">FooListParams</a>) (<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go/packages/pagination#PageNumber">PageNumber</a>[<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go">brucetestapi</a>.<a href="https://pkg.go.dev/github.com/DefinitelyATestOrg/test-api-go#FooListResponse">FooListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
