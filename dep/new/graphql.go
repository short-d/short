package new

import (
	"short/fw"
	"short/modern/mdgraphql"
)

type GraphQlPath string

func GraphGophers(graphqlPath GraphQlPath, logger fw.Logger, tracer fw.Tracer, g fw.GraphQlApi) fw.Server {
	return mdgraphql.NewGraphGophers(string(graphqlPath), logger, tracer, g)
}
