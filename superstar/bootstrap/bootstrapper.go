package bootstrap

import (
	"github.com/kataras/iris"
)

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./public/"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName   string
	AppAuthor string
}

func New(appName, AppAuthor string, configs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:     appName,
		AppAuthor:   AppAuthor,
		Application: iris.New(),
	}

	for _, cfg := range configs {
		cfg(b)
	}

	return b
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./views", ".html")

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	return b
}

// SetupViews loads the templates.
func (b *Bootstrapper) SetupViews(viewsDir, extension string) {
	htmlEngine := iris.HTML(viewsDir, extension).Layout("shared/layout.html")
	// 每次重新加载模版（线上关闭它）
	//htmlEngine.Reload(false)

	b.RegisterView(htmlEngine)
}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
