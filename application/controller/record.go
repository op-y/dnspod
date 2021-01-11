package controller

import (
	"log"
	"net/http"

	"github.com/op-y/dnspod/config"
	"github.com/op-y/dnspod/dnspod"

	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

type QueryRecordInput struct {
	DomainID  string `form:"domain_id" json:"domain_id"`
	Domain    string `form:"domain" json:"domain"`
	Offset    string `form:"offset" json:"offset"`
	Length    string `form:"length" json:"length"`
	SubDomain string `form:"sub_domain" json:"sub_domain"`
	Keyword   string `form:"keyword" json:"keyword"`
}

type CreateRecordInput struct {
	DomainID     string `form:"domain_id" json:"domain_id"`
	Domain       string `form:"domain" json:"domain"`
	SubDomain    string `form:"sub_domain" json:"sub_domain"`
	RecordType   string `form:"record_type" json:"record_type"`
	RecordLine   string `form:"record_line" json:"record_line"`
	RecordLineID string `form:"record_line_id" json:"record_line_id"`
	Value        string `form:"value" json:"value"`
	MX           string `form:"mx" json:"mx"`
	TTL          string `form:"ttl" json:"ttl"`
	Status       string `form:"status" json:"status"`
	Weight       string `form:"weight" json:"weight"`
}

type ModifyRecordInput struct {
	DomainID     string `form:"domain_id" json:"domain_id"`
	Domain       string `form:"domain" json:"domain"`
	RecordID     string `form:"record_id" json:"record_id"`
	SubDomain    string `form:"sub_domain" json:"sub_domain"`
	RecordType   string `form:"record_type" json:"record_type"`
	RecordLine   string `form:"record_line" json:"record_line"`
	RecordLineID string `form:"record_line_id" json:"record_line_id"`
	Value        string `form:"value" json:"value"`
	MX           string `form:"mx" json:"mx"`
	TTL          string `form:"ttl" json:"ttl"`
	Status       string `form:"status" json:"status"`
	Weight       string `form:"weight" json:"weight"`
}

type QueryRecordInfoInput struct {
	DomainID string `form:"domain_id" json:"domain_id"`
	Domain   string `form:"domain" json:"domain"`
	RecordID string `form:"record_id" json:"record_id"`
}

type SetRecordStatusInput struct {
	DomainID string `form:"domain_id" json:"domain_id"`
	Domain   string `form:"domain" json:"domain"`
	RecordID string `form:"record_id" json:"record_id"`
	Status   string `form:"status" json:"status"`
}

func CreateRecord(c *gin.Context) {
	var input CreateRecordInput
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
		output.Message = "domain_id and domain can not be empty at the same time."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.SubDomain {
		param.Add("sub_domain", input.SubDomain)
	}

	if "" != input.RecordType {
		param.Add("record_type", input.RecordType)
	}

	if "" != input.RecordLineID {
		param.Add("record_line_id", input.RecordLineID)
	} else if "" != input.RecordLine {
		param.Add("record_line", input.RecordLine)
	} else {
		output.Code = http.StatusBadRequest
		output.Message = "record line id and record line can not be empty at the same time."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.Value {
		param.Add("value", input.Value)
	} else {
		output.Code = http.StatusBadRequest
		output.Message = "value should not be empty."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.MX {
		param.Add("mx", input.MX)
	}

	if "" != input.TTL {
		param.Add("ttl", input.TTL)
	}

	if "" != input.Status {
		param.Add("status", input.Status)
	}

	if "" != input.Weight {
		param.Add("weight", input.Weight)
	}

	param.ToSlice()
	paramData := param.ToString()

	url := "https://dnsapi.cn/Record.Create"

	result, err := dnspod.CallAPI(url, config.CFG.Token, paramData)
	if err != nil {
		log.Printf("fail to call DNSPod API: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to create record"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	json, err := simplejson.NewJson(result)
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to create record"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	code, err := json.Get("status").Get("code").String()
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to create record"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	if "1" != code {
		log.Printf("response code is not 1")
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get create record"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	output.Code = http.StatusOK
	output.Message = "create record successfully"
	output.Data = string(result)
	c.JSON(http.StatusOK, output)
	return
}

func GetRecord(c *gin.Context) {
	var input QueryRecordInput
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
		output.Message = "domain_id and domain can not be empty at the same time."
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

	if "" != input.SubDomain {
		param.Add("sub_domain", input.SubDomain)
	}

	if "" != input.Keyword {
		param.Add("keyword", input.Keyword)
	}

	param.ToSlice()
	paramData := param.ToString()

	url := "https://dnsapi.cn/Record.List"

	result, err := dnspod.CallAPI(url, config.CFG.Token, paramData)
	if err != nil {
		log.Printf("fail to call DNSPod API: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get record list"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	json, err := simplejson.NewJson(result)
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get record list"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	code, err := json.Get("status").Get("code").String()
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get record list"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	if "1" != code {
		log.Printf("response code is not 1")
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get record list"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	output.Code = http.StatusOK
	output.Message = "get record list successfully"
	output.Data = string(result)
	c.JSON(http.StatusOK, output)
	return
}

func ModifyRecord(c *gin.Context) {
	var input ModifyRecordInput
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
		output.Message = "domain_id and domain can not be empty at the same time."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.RecordID {
		param.Add("record_id", input.RecordID)
	} else {
		output.Code = http.StatusBadRequest
		output.Message = "record ID should not be empty."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.SubDomain {
		param.Add("sub_domain", input.SubDomain)
	}

	if "" != input.RecordType {
		param.Add("record_type", input.RecordType)
	}

	if "" != input.RecordLineID {
		param.Add("record_line_id", input.RecordLineID)
	} else if "" != input.RecordLine {
		param.Add("record_line", input.RecordLine)
	} else {
		output.Code = http.StatusBadRequest
		output.Message = "record line id and record line can not be empty at the same time."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.Value {
		param.Add("value", input.Value)
	} else {
		output.Code = http.StatusBadRequest
		output.Message = "value should not be empty."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.MX {
		param.Add("mx", input.MX)
	}

	if "" != input.TTL {
		param.Add("ttl", input.TTL)
	}

	if "" != input.Status {
		param.Add("status", input.Status)
	}

	if "" != input.Weight {
		param.Add("weight", input.Weight)
	}

	param.ToSlice()
	paramData := param.ToString()

	url := "https://dnsapi.cn/Record.Modify"

	result, err := dnspod.CallAPI(url, config.CFG.Token, paramData)
	if err != nil {
		log.Printf("fail to call DNSPod API: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to modify record"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	json, err := simplejson.NewJson(result)
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to modify record"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	code, err := json.Get("status").Get("code").String()
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to modify record"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	if "1" != code {
		log.Printf("response code is not 1")
		output.Code = http.StatusInternalServerError
		output.Message = "fail to modify record"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	output.Code = http.StatusOK
	output.Message = "modify record successfully"
	output.Data = string(result)
	c.JSON(http.StatusOK, output)
}

func GetRecordInfo(c *gin.Context) {
	var input QueryRecordInfoInput
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
		output.Message = "domain_id and domain can not be empty at the same time."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.RecordID {
		param.Add("record_id", input.RecordID)
	} else {
		output.Code = http.StatusBadRequest
		output.Message = "record ID should not be empty."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	param.ToSlice()
	paramData := param.ToString()

	url := "https://dnsapi.cn/Record.Info"

	result, err := dnspod.CallAPI(url, config.CFG.Token, paramData)
	if err != nil {
		log.Printf("fail to call DNSPod API: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get record info"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	json, err := simplejson.NewJson(result)
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get record info"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	code, err := json.Get("status").Get("code").String()
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get record info"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	if "1" != code {
		log.Printf("response code is not 1")
		output.Code = http.StatusInternalServerError
		output.Message = "fail to get record info"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	output.Code = http.StatusOK
	output.Message = "get record info successfully"
	output.Data = string(result)
	c.JSON(http.StatusOK, output)
	return
}

func SetRecordStatus(c *gin.Context) {
	var input SetRecordStatusInput
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
		output.Message = "domain_id and domain can not be empty at the same time."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.RecordID {
		param.Add("record_id", input.RecordID)
	} else {
		output.Code = http.StatusBadRequest
		output.Message = "record ID should not be empty."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	if "" != input.Status {
		param.Add("status", input.Status)
	} else {
		output.Code = http.StatusBadRequest
		output.Message = "status not be empty."
		output.Data = ""
		c.JSON(http.StatusBadRequest, output)
		return
	}

	param.ToSlice()
	paramData := param.ToString()

	url := "https://dnsapi.cn/Record.Status"

	result, err := dnspod.CallAPI(url, config.CFG.Token, paramData)
	if err != nil {
		log.Printf("fail to call DNSPod API: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to set record status"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	json, err := simplejson.NewJson(result)
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to set record status"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	code, err := json.Get("status").Get("code").String()
	if err != nil {
		log.Printf("fail to parser response body: %s", err.Error())
		output.Code = http.StatusInternalServerError
		output.Message = "fail to set record status"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}
	if "1" != code {
		log.Printf("response code is not 1")
		output.Code = http.StatusInternalServerError
		output.Message = "fail to set record status"
		output.Data = ""
		c.JSON(http.StatusInternalServerError, output)
		return
	}

	output.Code = http.StatusOK
	output.Message = "set record status successfully"
	output.Data = string(result)
	c.JSON(http.StatusOK, output)
	return
}
