package resolver

type Query struct {
}

type UrlArgs struct {
	Id string
}

func (q Query) Url(args *UrlArgs) *Url {
	return &Url{}
}
