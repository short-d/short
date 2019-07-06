package fw

type Resolver = interface{}

type GraphQlApi interface {
	GetSchema() string
	GetResolver() Resolver
}
