package gogg

import (
	"net/http"
)

type Application struct {
	Name string
}

func NewApplication(name string) *Application {
	return &Application{
		Name: name,
	}
}

func (app *Application) Run() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>` + app.Name + `</title>
				</head>
				<body>
					<h1>` + app.Name + `</h1>
				</body>
			</html>
		`))
	})

	println("Servidor iniciado na porta 4000")
	http.ListenAndServe(":4000", nil)
}
