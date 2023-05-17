package httpd

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/transcriptaze/snyth/snythd/html"
	"github.com/transcriptaze/snyth/snythd/log"
)

type dispatcher struct {
	fs        fs.FS
	templates *template.Template
	debug     bool
}

var templates *template.Template

var folders = []string{"/pages/*.html"}

const GzipMinimum = 16384

func Run(address string, port uint, debug bool) {
	infof("HTTPD", "server initialising")

	// ... initialise HTML templates

	fs := html.HTML
	templates, err := template.New("snythd").Funcs(funcs).ParseFS(fs, "index.html")
	if err != nil {
		fatalf("HTTPD", "%v", err)
	}

	// ... initialise HTTP server
	httpd := dispatcher{
		fs:        fs,
		templates: templates,
		debug:     debug,
	}

	fsys := httpdFS{
		http.FS(fs),
	}

	static := cors(http.FileServer(fsys), debug)

	http.Handle("/css/", static)
	http.Handle("/fonts/", static)
	http.Handle("/images/", static)
	http.Handle("/javascript/", static)
	http.Handle("/midi/", static)
	http.Handle("/favicon.ico", static)
	http.Handle("/index.html", static)
	http.HandleFunc("/", httpd.dispatch)

	if p := os.Getenv("PORT"); p != "" {
		if u, err := strconv.ParseUint(p, 10, 32); err == nil {
			port = uint(u)
		}
	}

	infof("HTTPD", "server listening on address %v:%v", address, port)
	fatalf("HTTPD", "%v", http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), nil))
}

func (d *dispatcher) dispatch(w http.ResponseWriter, r *http.Request) {
	if url, err := url.QueryUnescape(fmt.Sprintf("%v", r.URL)); err == nil {
		debugf("HTTPD", "%-4v %v", r.Method, url)
	} else {
		debugf("HTTPD", "%-4v %v", r.Method, r.URL)
	}

	switch strings.ToUpper(r.Method) {
	case http.MethodOptions:
		d.options(w, r)

	case http.MethodGet:
		d.get(w, r)

	default:
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
	}
}

func cors(next http.Handler, debug bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")

		// devmode ?
		if debug {
			w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		} else {
			w.Header().Set("Cache-Control", "max-age=3600, must-revalidate")
		}

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func resolve(u *url.URL) (string, error) {
	base, err := url.Parse("/")
	if err != nil {
		return "", err
	}

	return base.ResolveReference(u).EscapedPath(), nil
}

func acceptsGzip(r *http.Request) bool {
	for k, h := range r.Header {
		if strings.TrimSpace(strings.ToLower(k)) == "accept-encoding" {
			for _, v := range h {
				if strings.Contains(strings.TrimSpace(strings.ToLower(v)), "gzip") {
					return true
				}
			}
		}
	}

	return false
}

func debugf(tag string, format string, values ...interface{}) {
	f := fmt.Sprintf("%-5v %v", tag, format)

	log.Debugf(f, values...)
}

func infof(tag string, format string, values ...interface{}) {
	f := fmt.Sprintf("%-5v %v", tag, format)

	log.Infof(f, values...)
}

func warnf(tag string, format string, values ...interface{}) {
	f := fmt.Sprintf("%-5v %v", tag, format)

	log.Warnf(f, values...)
}

func errorf(tag string, format string, values ...interface{}) {
	f := fmt.Sprintf("%-5v %v", tag, format)

	log.Errorf(f, values...)
}

func fatalf(tag string, format string, values ...interface{}) {
	f := fmt.Sprintf("%-5v %v", tag, format)

	log.Fatalf(f, values...)
}
