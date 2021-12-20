package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atrariksa/fastrogos/fenuma/configs"
	"github.com/atrariksa/fastrogos/fenuma/constants"
	"github.com/atrariksa/fastrogos/fenuma/handlers"
	"github.com/atrariksa/fastrogos/fenuma/middlewares"
	"github.com/atrariksa/fastrogos/fenuma/utils"
	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg := configs.Get()

	log := utils.GetLogger(cfg)

	// The HTTP Server
	server := &http.Server{Addr: cfg.App.Hostname, Handler: setupApis(cfg, log)}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx,
			time.Duration(cfg.App.Shutdown.GracePeriodSeconds)*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

var routes = flag.Bool("routes", true, "Generate router documentation")

func setupApis(cfg *configs.Config, logger *logrus.Logger) http.Handler {

	r := chi.NewRouter()
	loggerHandler := utils.NewStructuredLogger(logger)

	r.Use(chi_middleware.RequestID)
	r.Use(loggerHandler)
	r.Use(chi_middleware.Recoverer)
	r.Use(middlewares.DefaultResponseHeadersHandler(cfg))
	r.Use(cors.Handler(utils.GetCorsOptions(cfg)))

	handlers.WireHandlers(r, cfg, logger)

	if *routes {
		fmt.Println(docgen.JSONRoutesDoc(r))
	}

	return r
}

func writeDocsJSON() {
	cfg := configs.Get()

	dataByte, err := os.ReadFile("./docs/swagger.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	m := make(map[string]interface{})
	json.Unmarshal(dataByte, &m)

	sInfo := m["info"]
	inf := sInfo.(map[string]interface{})
	inf["title"] = constants.Fenuma
	inf["version"] = constants.Version
	m["host"] = cfg.App.Hostname
	m["basePath"] = "/"

	of, err := os.Create("./docs/" + constants.Fenuma + ".json")
	if err != nil {
		fmt.Println(err)
		return
	}
	mByte, err := json.Marshal(&m)
	of.Write(mByte)
}
