package main

import (
	"_iris/superstar/bootstrap"
	"_iris/superstar/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("superstar", "app-author")
	app.Bootstrap()
	app.Configure(routes.Configure)

	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
