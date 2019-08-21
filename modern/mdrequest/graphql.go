package mdrequest

import (
	"encoding/json"
	"net/http"
	"short/fw"
)

type graphQlResponse struct {
	Data interface{} `json:"data"`
}

type GraphQl struct {
	http fw.HTTPRequest
	root string
}

func (g GraphQl) Query(query fw.GraphQlQuery, headers map[string]string, response interface{}) error {
	var res graphQlResponse

	reqBuf, err := json.Marshal(query)
	if err != nil {
		return err
	}

	err = g.http.JSON(http.MethodPost, g.root, headers, string(reqBuf), &res)
	if err != nil {
		return err
	}

	resBuf, err := json.Marshal(res.Data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBuf, &response)
	return err
}

func (g GraphQl) RootUrl(root string) fw.GraphQlRequest {
	g.root = root
	return g
}

func NewGraphQl(http fw.HTTPRequest) fw.GraphQlRequest {
	return GraphQl{
		http: http,
	}
}
