package controller

import (
	"net/http"

	"github.com/op-y/dnspod/config"

	"github.com/gin-gonic/gin"
)

type OpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func StartGin(port string, r *gin.Engine) {
	SystemRoutes(r)
	AppRoutes(r)
	r.Run(port)
}

func SystemRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, DNSPod!")
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
}

func AppRoutes(r *gin.Engine) {
	dnsapi := r.Group("/v1", gin.BasicAuth(gin.Accounts{config.CFG.Username: config.CFG.Password}))

	dnsapi.GET("/user", GetUserDetail)
	dnsapi.GET("/user/log", GetUserLog)

	dnsapi.GET("/domain", GetDomainList)
	dnsapi.GET("/domain/info", GetDomainInfo)
	dnsapi.GET("/domain/log", GetDomainLog)

	dnsapi.POST("/record", CreateRecord)
	dnsapi.GET("/record", GetRecord)
	dnsapi.PUT("/record", ModifyRecord)
	dnsapi.GET("/record/info", GetRecordInfo)
	dnsapi.PUT("/record/status", SetRecordStatus)
}
