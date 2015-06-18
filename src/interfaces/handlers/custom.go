package handlers

// 26.04.15 13:41
// (c) Dmitriy Blokhin (sv.dblokhin@gmail.com), www.webjinn.ru

import (
    "interfaces/app"
)

type HandleCustom struct {
    app.HTTPController
}

func (h HandleCustom) GET(app *app.ContextApplication) {

    doc := app.Doc.Clone("custom") // like child
    doc["TitlePage"] = "Some text value (string)"

    result := doc.Compile()
    app.Ctx.SendHTML(result)
}

func (h HandleCustom) POST(app *app.ContextApplication) {
	// and so on...
}