package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

func (app *application) langToCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		langURL := chi.URLParam(r, "lang")
		ctx := context.WithValue(r.Context(), contextKeyLang, langURL)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) redirectLang(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := app.lang(r)
		acceptLang := r.Header.Get("Accept-Language")
		redirectTag, _ := app.localeManager.Match(acceptLang, lang).Base()
		redirectLang := redirectTag.String()

		if lang == redirectLang {
			next.ServeHTTP(w, r)
			return
		}

		var url string
		if strings.Contains(r.URL.Path, "/"+lang) {
			url = strings.Replace(r.URL.Path, "/"+lang, "/"+redirectLang, 1)
		} else {
			url = "/" + redirectLang + r.URL.Path
		}

		http.Redirect(w, r, url, http.StatusSeeOther)
	})
}

func (app *application) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				app.serverError(w, fmt.Errorf("%s", err))
				w.Header().Set("Connection", "close")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
