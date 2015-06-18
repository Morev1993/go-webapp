package handlers

// 18.02.15 0:44
// (c) Dmitriy Blokhin (sv.dblokhin@gmail.com), www.webjinn.ru

import (
	"interfaces/app"
	"net/http"
)

type Handle404 struct {
	app.HTTPController
}

func (h Handle404) GET(app *app.ContextApplication) {
	doc := app.Doc.Clone("404")
	page := doc.Compile()

	app.Ctx.Output.SetStatus(http.StatusNotFound)
	app.Ctx.SendHTML(page)
}
