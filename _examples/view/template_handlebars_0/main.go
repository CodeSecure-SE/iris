package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	// Init the handlebars engine
	engine := iris.Handlebars("./templates", ".html").Reload(true)
	// Register a helper.
	engine.AddFunc("fullName", func(person map[string]string) string {
		return person["firstName"] + " " + person["lastName"]
	})

	app.RegisterView(engine)

	app.Get("/", func(ctx iris.Context) {
		viewData := iris.Map{
			"author": map[string]string{"firstName": "Jean", "lastName": "Valjean"},
			"body":   "Life is difficult",
			"comments": []iris.Map{{
				"author": map[string]string{"firstName": "Marcel", "lastName": "Beliveau"},
				"body":   "LOL!",
			}},
		}

		ctx.View("example.html", viewData)
	})

	// Read more about its syntax at:
	// https://github.com/aymerick/raymond and
	// https://handlebarsjs.com/guide
	app.Listen(":8080")
}