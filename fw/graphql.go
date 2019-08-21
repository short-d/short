package fw

type Resolver = interface{}

type GraphQlAPI interface {
	GetSchema() string
	GetResolver() Resolver
}

type Scalar interface {
	ImplementsGraphQLType(name string) bool
	UnmarshalGraphQL(input interface{}) error
	MarshalJSON() ([]byte, error)
}
