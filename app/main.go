package main

import (
	"app/src"
	"app/src/metrics"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-logr/stdr"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"log"
	"os"
)

func main() {
	stdr.SetVerbosity(5)
	initConfig()

	cleanup := metrics.InitTracer()
	defer cleanup(context.Background())

	r := setupRouter()
	r.Use(otelgin.Middleware(metrics.ServiceName))
	r.Use(metrics.PrometheusMiddleware)
	_ = r.Run(":8000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	userRepo := src.CreateNewEntitiesRepo()

	r.GET("/entities", userRepo.GetEntities)
	r.GET("/entities/:id", userRepo.GetEntity)
	r.DELETE("/entities/:id", userRepo.DeleteEntity)
	r.GET("/uploadEntities", userRepo.UploadEntitiesToDb)
	r.DELETE("/entities", userRepo.DeleteEntities)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return r
}

func initConfig() {
	if os.Getenv("ENVIRONMENT") != "k8s" {
		err := godotenv.Load("local.env")
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
