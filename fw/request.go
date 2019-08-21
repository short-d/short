package fw

type HTTPRequest interface {
	JSON(method string, url string, headers map[string]string, body string, v interface{}) error
}

type GraphQlQuery struct {
	Query     string            `json:"query"`
	Variables map[string]string `json:"variables"`
}

type GraphQlRequest interface {
	RootUrl(root string) GraphQlRequest
	Query(request GraphQlQuery, headers map[string]string, response interface{}) error
}
