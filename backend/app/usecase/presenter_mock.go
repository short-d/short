package usecase

type showHomeCallArgs struct {
	authToken string
}

type showUserHomeCallArgs struct {
	authToken string
}

type showExternalPageCallArgs struct {
	link string
}

var _ Presenter = (*mockPresenter)(nil)

type mockPresenter struct {
	showHomeCallArgs         []showHomeCallArgs
	showUserHomeCallArgs     []showUserHomeCallArgs
	showExternalPageCallArgs []showExternalPageCallArgs
}

func (m *mockPresenter) ShowHome() {
	m.showHomeCallArgs = append(m.showHomeCallArgs, showHomeCallArgs{})
}

func (m *mockPresenter) ShowUserHome(authToken string) {
	m.showUserHomeCallArgs = append(
		m.showUserHomeCallArgs,
		showUserHomeCallArgs{authToken: authToken},
	)
}

func (m mockPresenter) ShowAliasNotFound() {
	panic("implement me")
}

func (m mockPresenter) ShowLongLink(longLink string) {
	panic("implement me")
}

func (m *mockPresenter) ShowExternalPage(link string) {
	m.showExternalPageCallArgs = append(
		m.showExternalPageCallArgs,
		showExternalPageCallArgs{link: link},
	)
}

func (m mockPresenter) ShowInvalidAuthTokenError() {
	panic("implement me")
}

func newMockPresenter() mockPresenter {
	return mockPresenter{
		showHomeCallArgs:         make([]showHomeCallArgs, 0),
		showUserHomeCallArgs:     make([]showUserHomeCallArgs, 0),
		showExternalPageCallArgs: make([]showExternalPageCallArgs, 0),
	}
}
