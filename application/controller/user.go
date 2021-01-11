package controller

import (
	"log"
	"net/http"

	"github.com/op-y/dnspod/config"
	"github.com/op-y/dnspod/dnspod"

	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

func GetUserDetail(c *gin.Context) {
	var output OpResponse

	url := "https://dnsapi.cn/User.Detail"

	// call API
	result, err := dnspod.CallAPI(url, config.CFG.Token, "")
	if err != nil {
		log.Printf("fail to call DNSPod API: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get user detail"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	json, err := simplejson.NewJson(result)
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get user detail"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	code, err := json.Get("status").Get("code").String()
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get user detail"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	if "1" != code {
		log.Printf("response code is not 1")
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get user detail"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	output.Code = http.StatusOK
	output.Message = "get user detail successfully"
	output.Data = string(result)
	c.JSON(http.StatusOK, output)
	return
}

func GetUserLog(c *gin.Context) {
	var output OpResponse

	url := "https://dnsapi.cn/User.Log"

	result, err := dnspod.CallAPI(url, config.CFG.Token, "")
	if err != nil {
		log.Printf("fail to call DNSPod API: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get user logs"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	json, err := simplejson.NewJson(result)
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get user logs"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	code, err := json.Get("status").Get("code").String()
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get user logs"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	if "1" != code {
		log.Printf("response code is not 1")
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get user logs"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	output.Code = http.StatusOK
	output.Message = "get user logs successfully"
	output.Data = string(result)
	c.JSON(http.StatusOK, output)
	return
}
