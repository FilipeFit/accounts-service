package app

import (
	"fmt"
	ginopentracing "github.com/Bose/go-gin-opentracing"
	"github.com/filipeFit/account-service/config"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"io"
	"log"
	"os"
)

var (
	router *gin.Engine
)

func init() {
	gin.SetMode(gin.DebugMode)
	router = gin.Default()
}

func StartApp() {
	handleTracing()
	mapRoutes()
	appPort := fmt.Sprintf(":%s", config.Config.AppPort)
	if err := router.Run(appPort); err != nil {
		log.Fatal(err)
	}
}

func handleTracing() {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "unknown"
	}
	// initialize the global singleton for tracing...
	tracer, reporter, closer, err := ginopentracing.InitTracing(
		fmt.Sprintf("account-service::%s", hostName),
		"localhost:5775",
		ginopentracing.WithEnableInfoLog(true))

	if err != nil {
		log.Fatal(fmt.Sprintf("error initializing tracer :%s", err))
	}
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {

		}
	}(closer)
	defer reporter.Close()
	opentracing.SetGlobalTracer(tracer)

	p := ginopentracing.OpenTracer([]byte("api-request-"))
	router.Use(p)
}
