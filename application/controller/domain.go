package controller

import (
	"log"
	"net/http"

	"github.com/op-y/dnspod/config"
	"github.com/op-y/dnspod/dnspod"

	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

type QueryDomainInput struct {
	Type    string `form:"type" json:"type"`
	Offset  string `form:"offset" json:"offset"`
	Length  string `form:"length" json:"length"`
	GroupID string `form:"group_id" json:"group_id"`
	Keyword string `form:"keyword" json:"keyword"`
}

type QueryDomainInfoInput struct {
	Domain   string `form:"domain" json:"domain"`
	DomainID string `form:"domain_id" json:"domain_id"`
}

type QueryDomainLogInput struct {
	Domain   string `form:"domain" json:"domain"`
	DomainID string `form:"domain_id" json:"domain_id"`
	Offset   string `form:"offset" json:"offset"`
	Length   string `form:"length" json:"length"`
}

func GetDomainList(c *gin.Context) {
	var input QueryDomainInput
	var output OpResponse

	param := dnspod.NewParam()

	if c.ShouldBind(&input) != nil {
		output.Code = http.StatusBadRequest
		output.Message = "failed to bind parameters"
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.Type {
		param.Add("type", input.Type)
	}

	if "" != input.Offset {
		param.Add("offset", input.Offset)
	}

	if "" != input.Length {
		param.Add("length", input.Length)
	}

	if "" != input.GroupID {
		param.Add("group_id", input.GroupID)
	}

	if "" != input.Keyword {
		param.Add("keyword", input.Keyword)
	}

	param.ToSlice()
	paramData := param.ToString()

	url := "https://dnsapi.cn/Domain.List"

	result, err := dnspod.CallAPI(url, config.CFG.Token, paramData)
	if err != nil {
		log.Printf("fail to call DNSPod API: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain list"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	json, err := simplejson.NewJson(result)
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain list"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	code, err := json.Get("status").Get("code").String()
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain list"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	if "1" != code {
		log.Printf("response code is not 1")
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain list"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	output.Code = http.StatusOK
	output.Message = "get domain list successfully"
	output.Data = string(result)
	c.JSON(http.StatusOK, output)
	return
}

func GetDomainInfo(c *gin.Context) {
	var input QueryDomainInfoInput
	var output OpResponse

	param := dnspod.NewParam()

	if c.ShouldBind(&input) != nil {
		output.Code = http.StatusBadRequest
		output.Message = "failed to bind parameters"
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.DomainID {
		param.Add("domain_id", input.DomainID)
	} else if "" != input.Domain {
		param.Add("domain", input.Domain)
	} else {
		output.Code = http.StatusBadRequest
		output.Message = "domain_id and domain is empty at the same time."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	param.ToSlice()
	paramData := param.ToString()

	url := "https://dnsapi.cn/Domain.Info"

	result, err := dnspod.CallAPI(url, config.CFG.Token, paramData)
	if err != nil {
		log.Printf("fail to call DNSPod API: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain info"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	json, err := simplejson.NewJson(result)
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain info"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	code, err := json.Get("status").Get("code").String()
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain info"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	if "1" != code {
		log.Printf("response code is not 1")
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain info"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	output.Code = http.StatusOK
	output.Message = "get domain info successfully"
	output.Data = string(result)
	c.JSON(http.StatusOK, output)
	return
}

func GetDomainLog(c *gin.Context) {
	var input QueryDomainLogInput
	var output OpResponse

	param := dnspod.NewParam()

	if c.ShouldBind(&input) != nil {
		output.Code = http.StatusBadRequest
		output.Message = "failed to bind parameters"
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.DomainID {
		param.Add("domain_id", input.DomainID)
	} else if "" != input.Domain {
		param.Add("domain", input.Domain)
	} else {
		output.Code = http.StatusBadRequest
		output.Message = "domain_id and domain is empty at the same time."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.Offset {
		param.Add("offset", input.Offset)
	}

	if "" != input.Length {
		param.Add("length", input.Length)
	}

	param.ToSlice()
	paramData := param.ToString()

	url := "https://dnsapi.cn/Domain.Log"

	result, err := dnspod.CallAPI(url, config.CFG.Token, paramData)
	if err != nil {
		log.Printf("fail to call DNSPod API: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain log"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	json, err := simplejson.NewJson(result)
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain log"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	code, err := json.Get("status").Get("code").String()
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get domain log"
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
	output.Message = "get domain log successfully"
	output.Data = string(result)
	c.JSON(http.StatusOK, output)
	return
}
