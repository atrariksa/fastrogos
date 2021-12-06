package main

//go:generate go run github.com/swaggo/swag/cmd/swag init
import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atrariksa/fastrogos/rula/configs"
	"github.com/atrariksa/fastrogos/rula/drivers"
	"github.com/atrariksa/fastrogos/rula/handlers"
	"github.com/atrariksa/fastrogos/rula/middlewares"
	"github.com/atrariksa/fastrogos/rula/migrations"
	"github.com/atrariksa/fastrogos/rula/utils"
	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func main() {

	cmdMessage :=
		`
	Please use following commands :
	1. use "migrate" to migrate tables
	2. use "server" to run service
	`
	if len(os.Args) == 1 {
		log.Fatalln(cmdMessage)
	}
	command := os.Args[1]
	switch command {
	case "server":
		server()
	case "migrate":
		migrate(os.Args)
	case "write_docs":
		writeDocsJSON()
	default:
		log.Println(fmt.Sprintf(`Unknown command "%v". %v`, command, cmdMessage))
	}
}

func server() {

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

	return r
}

func migrate([]string) {
	cfg := configs.Get()
	dbWrite := drivers.NewDBClient(cfg)
	m := migrations.Migrator{DB: dbWrite}
	m.MigrateUp()
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
	inf["title"] = "Rula"
	inf["version"] = "0.0.1"
	m["host"] = cfg.App.Hostname
	m["basePath"] = "/"

	of, err := os.Create("./docs/rula.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	mByte, err := json.Marshal(&m)
	of.Write(mByte)
}
