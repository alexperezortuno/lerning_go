package learning

import (
	"../../internal/api"
	"../../internal/util"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"time"
)

func Init() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	version := os.Getenv("VERSION")

	server := &http.Server{
		Addr:           host + ":" + port,
		Handler:        api.NewRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Configure Logging
	LogFileLocation := os.Getenv("LOG_FILE_LOCATION")

	if LogFileLocation != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LogFileLocation,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     5,    //days
			Compress:   true, // disabled by default
		})
	}

	log.Println("Starting Server, Version:", version)
	log.Println("Run server in PORT:", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	util.WaitForShutdown(server)
}
