package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/alecthomas/kong"
	"github.com/awryme/dnsproxy/pkg/dns/dnsserver"
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/rewrites/datastore/datastorejson"
	"github.com/awryme/dnsproxy/pkg/ui/uiserver"
	"github.com/awryme/slogf"
	"github.com/oklog/run"
)

var logOutput = os.Stdout

type App struct {
	RewritesCfg string `help:"path to config with rewrites" required:"" type:"existingfile"`
	Addr        string `help:"address to listen on (udp, http), without port" default:"127.0.0.1"`

	DNS struct {
		Port     int      `help:"dns server: port to listen on (udp)" default:"53"`
		Upstream *url.URL `help:"dns server: upstream dns url (schemas: tls)" required:""`
		BaseDns  string   `help:"dns server: base dns address for tls & https upstream resolve" default:"1.1.1.1"`
	} `embed:"" prefix:"dns."`
	UI struct {
		Port      int  `help:"ui server: port to listen on (http)" default:"8080"`
		LogAccess bool `help:"log ui api requests" default:"true"`
	} `embed:"" prefix:"ui."`
}

func (app *App) Run() error {
	ctx := context.Background()

	logHandler := slogf.DefaultHandler(logOutput)
	logf := slogf.New(logHandler)
	printBuildInfo(logf)

	datastore, err := datastorejson.New(logf, app.RewritesCfg)
	if err != nil {
		return fmt.Errorf("create rewrites json datastore: %w", err)
	}
	rstorage, err := rewrites.NewStorage(datastore, app.Addr)
	if err != nil {
		return fmt.Errorf("get rewrites cache storage: %w", err)
	}

	group := new(run.Group)
	onInterrupt := func(err error) {}

	runDnsServer := func() error {
		return dnsserver.Start(ctx, logf, dnsserver.Params{
			Addr:            app.Addr,
			Port:            app.DNS.Port,
			Upstream:        app.DNS.Upstream,
			BaseDns:         app.DNS.BaseDns,
			RewritesStorage: rstorage,
		})
	}

	runUIServer := func() error {
		settingsInfo := map[string]string{
			"self address":  app.Addr,
			"dns port":      fmt.Sprint(app.DNS.Port),
			"dns upstream":  app.DNS.Upstream.String(),
			"dns base addr": app.DNS.BaseDns,
			"ui port":       fmt.Sprint(app.UI.Port),
			"ui log access": fmt.Sprint(app.UI.LogAccess),
			"log output":    logOutput.Name(),
		}
		return uiserver.Start(ctx, logHandler, uiserver.Params{
			Addr:            app.Addr,
			Port:            app.UI.Port,
			RewritesStorage: rstorage,
			LogAccess:       app.UI.LogAccess,
			SettingsInfo:    settingsInfo,
		})
	}

	group.Add(runUIServer, onInterrupt)
	group.Add(runDnsServer, onInterrupt)

	return group.Run()
}

func main() {
	var app App
	kctx := kong.Parse(&app, kong.DefaultEnvars("DNSPROXY"))
	err := kctx.Run()
	kctx.FatalIfErrorf(err)
}
