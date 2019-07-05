package resolver

type Query struct {
}

type UrlAliasArgs struct {
	Id string
}

func (q Query) UrlAlias(args *UrlAliasArgs) *UrlAlias {
	return &UrlAlias{}
}
