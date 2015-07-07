# Разработка web-приложения на Golang

## Пример использования
Добавление собственных роутеров в `main.go`
```
Application.Routes(app.MapRoutes{
			{"/one": handlers.HandleOne{}},
			{"/two": handlers.HandleWwo{}},
			{"/{url:.*}": handlers.Handle404{}},
		})
```
Где каждый Handle должен быть создан в `src/interfaces/handlers/`
```
import (
	app "github.com/dblokhin/webapp"
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
```