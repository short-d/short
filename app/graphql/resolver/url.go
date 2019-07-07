package resolver

type Url struct {
}

func (u Url) Id() *string {
	id := "1"
	return &id
}

func (u Url) OriginalUrl() *string {
	url := "http://time4hacks.com"
	return &url
}

func (u Url) CustomAlias() *string {
	alias := "tiny"
	return &alias
}

func (u Url) ExpirationDate() *string {
	date := "01/01/95"
	return &date
}
