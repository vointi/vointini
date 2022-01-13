package server

import (
	"embed"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/vointi/vointini/frontend/templates"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
)

//go:embed templates/index.html
var mainTemplate embed.FS

func New() (router *chi.Mux) {
	endpoint := newApi()

	router = chi.NewRouter()
	router.Get(`/`, endpoint.index) // default redirect
	router.Get(`/global.css`, endpoint.global_css)
	router.Get(`/bootstrap.min.css`, endpoint.bootstrap_min_css)
	router.Get(`/favicon.png`, endpoint.favicon)

	router.Get(`/{page}.js`, endpoint.getPageJS)
	router.Get(`/{page}.js.map`, endpoint.getPageJSMap)
	router.Get(`/{page}.css`, endpoint.getPageCSS)

	router.Route(`/{language:[a-zA-Z-]{2,5}}`, func(r chi.Router) {
		r.Get(`/`, endpoint.index) // default redirect
		r.Get(`/{page}.html`, endpoint.getPage)
	})

	return router
}

type FrontEndServer struct {
	basepath string
}

func newApi() (fe *FrontEndServer) {
	fe = &FrontEndServer{}

	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	fe.basepath = cwd

	return fe
}

// index redirects to default language and main page
func (api FrontEndServer) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Location`, `/fi/entries.html`)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (api FrontEndServer) global_css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, `text/css`)

	fh, err := templates.Frontend.Open(path.Join(`public`, `global.css`))
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(w, fh)
	if err != nil {
		panic(err)
	}
}

// favicon returns FavIcon (PNG image)
func (api FrontEndServer) favicon(w http.ResponseWriter, r *http.Request) {
	fh, err := templates.Frontend.Open(path.Join(`public`, `favicon.png`))
	if err != nil {
		panic(err)
	}

	w.Header().Set(`Content-Type`, `image/png`)
	_, err = io.Copy(w, fh)
	if err != nil {
		panic(err)
	}
}

func (api FrontEndServer) bootstrap_min_css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, `text/css`)

	fh, err := templates.Frontend.Open(path.Join(`public`, `bootstrap.min.css`))
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(w, fh)
	if err != nil {
		panic(err)
	}
}

func (api FrontEndServer) getPageJSMap(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(`Content-Type`, `application/json`)

	page := chi.URLParam(request, `page`)

	fh, err := templates.Frontend.Open(path.Join(`public`, `build`, page, `bundle.js.map`))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		panic(err)
	}

	_, err = io.Copy(writer, fh)
	if err != nil {
		panic(err)
	}
}

func (api FrontEndServer) getPageJS(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(`Content-Type`, `application/javascript`)
	writer.Header().Set(`SourceMap`, request.RequestURI+`.map`)

	page := chi.URLParam(request, `page`)

	fh, err := templates.Frontend.Open(path.Join(`public`, `build`, page, `bundle.js`))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		panic(err)
	}

	_, err = io.Copy(writer, fh)
	if err != nil {
		panic(err)
	}
}

func (api FrontEndServer) getPageCSS(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(`Content-Type`, `text/css`)

	page := chi.URLParam(request, `page`)

	fh, err := templates.Frontend.Open(path.Join(`public`, `build`, page, `bundle.css`))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		panic(err)
	}

	_, err = io.Copy(writer, fh)
	if err != nil {
		panic(err)
	}
}

func (api FrontEndServer) getPage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(`Content-Type`, `text/html`)
	page := chi.URLParam(request, `page`)

	tpl, err := template.ParseFS(
		mainTemplate, path.Join(`templates`, `index.html`),
	)

	if err != nil {
		panic(err)
		return
	}

	err = tpl.Execute(writer, struct {
		Language string
		Page     string
	}{
		Language: `fi`,
		Page:     page,
	})

	if err != nil {
		panic(err)
	}
}
