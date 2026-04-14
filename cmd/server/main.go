package main

import (
	"flag"
	"fmt"
	"os"

	serverapp "go-template/internal/app/server"
	pkgconfig "go-template/pkg/config"
	pkgerrors "go-template/pkg/errors"
	pkglog "go-template/pkg/log"
	pkgversion "go-template/pkg/version"
)

func main() {
	var (
		configPath  string
		showVersion bool
	)

	flag.StringVar(&configPath, "config", "", "path to configurtion file")
	flag.BoolVar(&showVersion, "version", false, "print version information")
	flag.Parse()

	if showVersion {
		fmt.Println(pkgversion.String())
		os.Exit(0)
	}

	cfg := pkgerrors.Must(pkgconfig.Read[serverapp.Config](configPath))

	log := pkgerrors.Must(pkglog.NewLogger(
		pkglog.WithLevel(cfg.Log.Level),
		pkglog.WithOutputPaths(cfg.Log.OutputPaths...),
	))
	defer func() {
		_ = log.Sync()
	}()

	app, err := serverapp.NewApp(cfg, log)
	if err != nil {
		panic(err)
	}

	_ = app
}
