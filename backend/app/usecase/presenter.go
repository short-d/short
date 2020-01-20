package usecase

// Presenter transforms data into proper format and passes it to view layer.
type Presenter interface {
	ShowHome()
	ShowUserHome(authToken string)
	ShowAliasNotFound()
	ShowLongLink(longLink string)
	ShowExternalPage(link string)
	ShowInvalidAuthTokenError()
}
