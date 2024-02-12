package main

import (
	"net/http"

	"github.com/bmerchant22/hc_hackathon/auth"
	"github.com/bmerchant22/hc_hackathon/mail"
	"github.com/bmerchant22/hc_hackathon/middleware"
	"github.com/gin-gonic/gin"
)

func authServer(mail_channel chan mail.Mail) *http.Server {
	// PORT := viper.GetString("PORT.AUTH")
	PORT := "8080"
	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	auth.Router(mail_channel, r)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
