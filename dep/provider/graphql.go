package provider

import (
	"github.com/byliuyang/app/modern/mdgraphql"

	"github.com/byliuyang/app/fw"
)

type GraphQlPath string

func GraphGophers(graphqlPath GraphQlPath, logger fw.Logger, tracer fw.Tracer, g fw.GraphQlAPI) fw.Server {
	return mdgraphql.NewGraphGophers(string(graphqlPath), logger, tracer, g)
}
