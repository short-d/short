package fw

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type Resolver = interface{}

type GraphQl interface {
	GetSchema() string
	GetResolver() Resolver
}

type GraphGophers struct {
}

func (g GraphGophers) ListenAndServe(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func NewGraphGophers(g GraphQl) Server {
	schema := graphql.MustParseSchema(g.GetSchema(), g.GetResolver())
	http.Handle("/graphql", &relay.Handler{
		Schema: schema,
	})
	return &GraphGophers{}
}
