package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dimiro1/health"
	kitlog "github.com/go-kit/log"
	_ "github.com/go-sql-driver/mysql"
	goconfig "github.com/iglin/go-config"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sgp-processor-svc/internal/processReadBot"
	"sgp-processor-svc/internal/processReadBot/process"
	"syscall"
)

func Run() {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	port := config.GetString("server.port")

	var kitlogger kitlog.Logger
	kitlogger = kitlog.NewJSONLogger(os.Stderr)
	kitlogger = kitlog.With(kitlogger, "time", kitlog.DefaultTimestamp)

	mux := http.NewServeMux()
	errs := make(chan error, 2)
	////////////////////////////////////////////////////////////////////////
	////////////////////////CORS///////////////////////////////////////////
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	handlerCORS := cors.Handler(mux)

	db, err := sql.Open("mysql", getStrConnection())
	if err != nil {
		log.Fatalf("unable to open database connection %s", err.Error())
	}

	////////////////////////////////////////////BOT////////////////////////////////////////////
	var processorBotService processReadBot.IProcessReaderBot

	processorBotService = process.NewProcessReaderBot(context.Background(), kitlogger)

	if config.GetBool("queue-bot-process.activebot") {
		go processorBotService.ProcessReaderInitializer()
	}
	////////////////////////////////////////////BOT////////////////////////////////////////////

	mux.Handle("/health", health.NewHandler())

	go func() {
		kitlogger.Log("listening", "transport", "http", "address", port)
		errs <- http.ListenAndServe(":"+port, handlerCORS)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		db.Close()
	}()
	kitlogger.Log("terminated", <-errs)
}

func getStrConnection() string {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	host := config.GetString("datasource.host")
	user := config.GetString("datasource.user")
	pass := config.GetString("datasource.pass")
	dbname := config.GetString("datasource.dbname")
	strconn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", user, pass, host, dbname)
	return strconn
}
