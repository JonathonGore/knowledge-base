package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/JonathonGore/dots/yaml"
	"github.com/JonathonGore/knowledge-base/config"
	"github.com/JonathonGore/knowledge-base/handlers"
	_ "github.com/JonathonGore/knowledge-base/logging"
	"github.com/JonathonGore/knowledge-base/server"
	"github.com/JonathonGore/knowledge-base/session/managers"
	"github.com/JonathonGore/knowledge-base/storage"
	"github.com/JonathonGore/knowledge-base/storage/sql"
)

func getSQLConfig(conf config.Config) sql.Config {
	return (sql.Config{
		DatabaseName: conf.GetString("database.name"),
		Username:     conf.GetString("database.user"),
		Password:     conf.GetString("database.password"),
		Host:         conf.GetString("database.host"),
	})
}

func main() {
	confFile := flag.String("config", "config.yml", "specify the config file to use")
	flag.Parse()

	var api handlers.API
	var conf config.Config
	var d storage.Driver

	conf, err := yaml.New(*confFile)
	if err != nil {
		log.Fatalf("unable to parse configuration file: %v", err)
	}

	d, err = sql.New(getSQLConfig(conf))
	if err != nil {
		log.Fatalf("unable to create sql driver: %v", err)
	}

	sm, err := managers.NewSMManager("knowledge_base", 3600*24*365, d)
	if err != nil {
		log.Fatalf("unable to create session manager: %v", err)
	}

	api, err = handlers.New(d, sm)
	if err != nil {
		log.Fatalf("unable to create handler: %v", err)
	}

	s, err := server.New(api, sm)
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}

	srv := &http.Server{
		Addr:      fmt.Sprintf(":%v", conf.GetInt("server.port")),
		Handler:   s,
		TLSConfig: &tls.Config{},
	}

	log.Printf("Starting server over http on port: %v", conf.GetInt("server.port"))
	log.Fatal(srv.ListenAndServe())
}
