package dnspod

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func CallAPI(url, token, paramData string) ([]byte, error) {
	paramPublic := fmt.Sprintf("login_token=%s&format=json&lang=en&error_on_empty=no", token)
	param := paramPublic
	if "" != paramData {
		param = paramPublic + "&" + paramData
	}

	//fmt.Println("=====Param=====")
	//fmt.Println(param)

	// call API
	timeout := time.Second * 5
	client := &http.Client{Timeout: timeout}
	request, err := http.NewRequest("POST", url, strings.NewReader(param))
	if err != nil {
		log.Printf("fail to build request: %s", err.Error())
		return nil, err
	}
	request.Header.Add("UserAgent", "DNS API/0.0.1")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		log.Printf("fail to call DNSPod API %s: %s", url, err.Error())
		return nil, err
	}
	defer response.Body.Close()

	//fmt.Println("=====Body=====")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("fail to read response data: %s", err.Error())
		return nil, err
	}
	//fmt.Printf("response body:\n%s\n", string(body))

	//fmt.Println("=====Cookies=====")
	//cookies := response.Cookies()
	//for _, cookie := range cookies {
	//    fmt.Printf("%s: %s\n", cookie.Name, cookie.Value)
	//}

	return body, nil
}
