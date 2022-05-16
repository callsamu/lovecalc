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

func (app *application) detectLanguage(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		langURL := chi.URLParam(r, "lang")
		langHeader := r.Header.Get("Accept-Language")

		if langURL != langHeader {
			redirectURL := strings.Replace(r.URL.Path, "/"+langURL+"/", "/"+langHeader+"/", 1)
			http.Redirect(w, r, redirectURL, http.StatusSeeOther)
			return
		}

		_, ok := app.localizers[langURL]
		if !ok {
			app.notFound(w)
			return
		}

		ctx := context.WithValue(r.Context(), "lang", langURL)
		next.ServeHTTP(w, r.WithContext(ctx))
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
