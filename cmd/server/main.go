package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vointini/vointini/backend/restapi"
	"github.com/vointini/vointini/backend/serviceapi"
	"github.com/vointini/vointini/backend/storage/postgres"
	"github.com/vointini/vointini/backend/storage/storageimpl"
	"github.com/vointini/vointini/frontend/server"
	"github.com/vointini/vointini/pkg/meta"
	"golang.org/x/text/language"
	"io"
	"net/http"
	"os"
)

type PostgreSQL struct {
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Database string `json:"db"`
	Host     string `json:"host"`
	Port     uint16 `json:"port"` // 1-65535
}

type Config struct {
	Backend struct {
		Postgres *PostgreSQL `json:"postgres"`
	} `json:"backend"`
}

func loadConfig(configPath string) (*Config, error) {
	f, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var tmp Config
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		return nil, err
	}

	return &tmp, nil
}

func main() {
	configFile := flag.String(`config`, `config.json`, `Config file`)
	listenHost := flag.String(`host`, `127.0.0.1`, `IP to listen to`)
	listenPort := flag.Int(`port`, 8080, `Port to listen to`)

	flag.Parse()

	config, err := loadConfig(*configFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `couldn't load config: %v `, err)
		os.Exit(1)
	}

	// Backend storage for service
	var storage storageimpl.Storage

	if config.Backend.Postgres != nil {
		pgcfg := *config.Backend.Postgres
		storage = postgres.New(pgcfg.User, pgcfg.Pass, pgcfg.Database, pgcfg.Host, pgcfg.Port)
	}

	if storage == nil {
		_, _ = fmt.Fprintf(os.Stderr, `no storage loaded`)
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	svc := serviceapi.New(storage, language.Finnish)

	frontendServer := server.New()
	apiServer := restapi.New(svc)

	// Add API to router
	router.Mount(`/api/v1`, apiServer)

	// Add front end page(s), CSS, etc. to router
	router.Mount(`/`, frontendServer)

	address := fmt.Sprintf(`%s:%d`, *listenHost, *listenPort)

	_, _ = fmt.Fprintf(os.Stdout, `Vointini %s HTTP server listening at http://%s/`+"\n", meta.VERSION, address)

	if err := http.ListenAndServe(address, router); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `error: %v`, err)
		os.Exit(1)
	}

}
