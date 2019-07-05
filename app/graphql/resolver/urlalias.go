package resolver

type UrlAlias struct {
}

func (u UrlAlias) Id() *string {
	id := "1"
	return &id
}

func (u UrlAlias) OriginalUrl() *string {
	url := "http://time4hacks.com"
	return &url
}

func (u UrlAlias) CustomAlias() *string {
	alias := "tiny"
	return &alias
}

func (u UrlAlias) ExpirationDate() *string {
	date := "01/01/95"
	return &date
}
