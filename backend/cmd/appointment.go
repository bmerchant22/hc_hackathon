package main

import (
	"net/http"

	"github.com/bmerchant22/hc_hackathon/appointment"
	_ "github.com/bmerchant22/hc_hackathon/config"
	"github.com/bmerchant22/hc_hackathon/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func appointmentServer() *http.Server {
	PORT := viper.GetString("PORT.APPOINTMENT")
	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	appointment.Router(r)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
