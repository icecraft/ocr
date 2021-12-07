package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/icecraft/ocr/pkg/config"
)

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func main() {
	configFile := flag.String("config_file", "/config/config.json", "config file used when program startup")
	port := flag.String("port", "8080", "web server port that used")
	flag.Parse()

	cfg := config.Config{}
	if err := config.LoadConfigFromFile(*configFile, &cfg); err != nil {
		log.Errorf("failed to load config, reason:%s", err.Error())
		os.Exit(-1)
	}
	if os.Getenv("DEBUG") != "" {
		log.Info("server config, details: %v", cfg)
	}

	// register route
	r := gin.Default()
	cors_config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Cookie"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Cookie"},
		MaxAge:           12 * time.Hour,
	}
	cors_config.AllowAllOrigins = true
	r.Use(cors.New(cors_config))
	r.GET("/ping", ping)

	r.Run(fmt.Sprintf(":%s", *port))
}
