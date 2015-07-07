package handlers

// 18.02.15 0:44
// (c) Dmitriy Blokhin (sv.dblokhin@gmail.com), www.webjinn.ru

import (
	app "github.com/dblokhin/webapp"
)

type HandleHome struct {
	app.HTTPController
}

func (h HandleHome) GET(app *app.ContextApplication) {
	doc := app.Doc.Clone("base")
	doc["content"] = "Hello world"
	page := doc.Compile()
	
	app.Ctx.SendHTML(page)
}
