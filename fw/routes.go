package fw

import "net/http"

type Params map[string]string

type Handler func(w http.ResponseWriter, r *http.Request, params Params)

type Routes map[string]Handler
