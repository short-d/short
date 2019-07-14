package fw

type Resolver = interface{}

type GraphQlApi interface {
	GetSchema() string
	GetResolver() Resolver
}

type Scalar interface {
	ImplementsGraphQLType(name string) bool
	UnmarshalGraphQL(input interface{}) error
	MarshalJSON() ([]byte, error)
}
