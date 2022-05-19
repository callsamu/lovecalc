package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/callsamu/lovecalc/pkg/cache"
	redisc "github.com/callsamu/lovecalc/pkg/cache/redis"
	"github.com/callsamu/lovecalc/pkg/core"
	"github.com/callsamu/lovecalc/pkg/translations"
	"github.com/go-redis/redis/v8"
	"golang.org/x/text/language"
)

type contextKey string

var (
	contextKeyLang = contextKey("lang")
	defaultLang    = language.English
)

type application struct {
	calculator    core.Calculator
	infoLog       *log.Logger
	errorLog      *log.Logger
	matchCache    cache.MatchCache
	templateCache map[string]*template.Template
	localeManager *LocaleManager
}

func newRedisClient(url string) (*redis.Client, error) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return redis.NewClient(opt), nil
}

func main() {
	addr := flag.Int("addr", 4000, "Port address to listen on")
	redisUrl := flag.String("redis_url", os.Getenv("REDIS_URL"), "Redis instance URL")
	algorithm := flag.String("algorithm", core.HashWithSHA1, "Algorithm to be used by the calculator")
	flag.Parse()

	c, err := core.NewCalculator(*algorithm)
	if err != nil {
		log.Fatal(err)
		return
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	rdb, err := newRedisClient(*redisUrl)
	if err != nil {
		log.Fatal(err)
		return
	}
	mc := redisc.NewMatchCache(rdb)

	bundle, err := translations.Load(translations.LocalesFS, defaultLang)
	if err != nil {
		log.Fatal(err)
	}

	lm := NewLocaleManager(bundle)
	tc, err := newTemplateCache("./ui/template/", lm)
	if err != nil {
		log.Fatal(err)
	}

	app := application{
		localeManager: lm,
		templateCache: tc,
		matchCache:    mc,
		calculator:    c,
		infoLog:       infoLog,
		errorLog:      errorLog,
	}
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", *addr),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.infoLog.Printf("Listening on port %d", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
