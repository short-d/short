package resolver

type Mutation struct {
}

type UrlAliasInput struct {
	OriginalUrl    *string
	CustomAlias    *string
	ExpirationDate *string
}

type CreateUrlAliasArgs struct {
	UrlAlias  *UrlAliasInput
	UserEmail *string
}

func (m Mutation) CreateUrlAlias(args *CreateUrlAliasArgs) *UrlAlias {
	return &UrlAlias{}
}
