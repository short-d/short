//+build wireinject

package dep

import (
	"github.com/google/wire"
	"tinyURL/app"
)

func InitializeApp() app.App {
	wire.Build(app.NewApp)
	return app.App{}
}
