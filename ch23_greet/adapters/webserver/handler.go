package webserver

import (
	"embed"
	// _ "embed"
	"html/template"
	"net/http"

	"ch23/greet/domain/interactions"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

var (
	//go:embed "markup/*"
	templates embed.FS
)

func NewHandler() (http.Handler, error) {
	tmpl, err := template.ParseFS(templates, "markup/*.gohtml")
	if err != nil {
		return nil, err
	}

	hdl := handler{tmpl: tmpl}
	mux := http.NewServeMux()
	mux.HandleFunc("/", hdl.form)
	mux.HandleFunc(greetPath, hdl.replyWith(interactions.Greet))
	mux.HandleFunc(cursePath, hdl.replyWith(interactions.Curse))
	return mux, nil
}

type handler struct {
	tmpl *template.Template
}

func (h *handler) replyWith(fn func(name string) string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.tmpl.ExecuteTemplate(w, "reply.gohtml", fn(r.Form.Get("name"))); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *handler) form(w http.ResponseWriter, _ *http.Request) {
	h.tmpl.ExecuteTemplate(w, "form.gohtml", nil)
}
