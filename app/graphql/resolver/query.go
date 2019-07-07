package resolver

type Query struct {
}

type UrlArgs struct {
	Alias          string
}

func (q Query) Url(args *UrlArgs) *Url {
	return &Url{}
}
