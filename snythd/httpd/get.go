package httpd

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/oxtoacart/bpool"
)

var funcs = template.FuncMap{}
var bufpool = bpool.NewBufferPool(32)

func (d dispatcher) get(w http.ResponseWriter, r *http.Request) {
	path, err := resolve(r.URL)
	if err != nil {
		http.Error(w, "invalid URL", http.StatusBadRequest)
		return
	}

	switch path {
	default:
		d.html("index.html", w, r)
	}
}

func (d dispatcher) html(page string, w http.ResponseWriter, r *http.Request) {
	infof("HTTPD", "GET  %v", r.URL.Path)

	templates := d.templates
	if d.debug {
		if t, err := template.New("snythd").Funcs(funcs).ParseFS(d.fs, "index.html"); err != nil {
			errorf("HTTPD", "%v", err)
			http.Error(w, fmt.Sprintf("Internal error (%v)", err), http.StatusNotFound)
			return
		} else {
			templates = t
		}
	}

	info := map[string]interface{}{}

	if err := d.render(w, templates, page, info); err != nil {
		errorf("HTTPD", "%v", err)
		http.Error(w, fmt.Sprintf("%v: internal error", page), http.StatusInternalServerError)
	}
}

func (d dispatcher) render(w http.ResponseWriter, templates *template.Template, page string, data map[string]interface{}) error {
	buffer := bufpool.Get()
	defer bufpool.Put(buffer)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")

	// devmode ?
	if d.debug {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	} else {
		w.Header().Set("Cache-Control", "max-age=3600, must-revalidate")
	}

	if t := templates.Lookup(page); t == nil {
		return fmt.Errorf("Missing template for %v", page)
	} else if err := t.Execute(buffer, data); err != nil {
		return err
	} else {
		buffer.WriteTo(w)
	}

	return nil
}
