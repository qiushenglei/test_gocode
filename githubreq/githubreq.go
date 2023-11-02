package githubreq

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v2"
)

type TResp struct {
	Args struct {
	} `json:"args"`
	Headers struct {
		Accept                  string `json:"Accept"`
		AcceptEncoding          string `json:"Accept-Encoding"`
		AcceptLanguage          string `json:"Accept-Language"`
		Host                    string `json:"Host"`
		SecChUa                 string `json:"Sec-Ch-Ua"`
		SecChUaMobile           string `json:"Sec-Ch-Ua-Mobile"`
		SecChUaPlatform         string `json:"Sec-Ch-Ua-Platform"`
		SecFetchDest            string `json:"Sec-Fetch-Dest"`
		SecFetchMode            string `json:"Sec-Fetch-Mode"`
		SecFetchSite            string `json:"Sec-Fetch-Site"`
		SecFetchUser            string `json:"Sec-Fetch-User"`
		UpgradeInsecureRequests string `json:"Upgrade-Insecure-Requests"`
		UserAgent               string `json:"User-Agent"`
		XAmznTraceId            string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

func Use() {
	var s TResp
	c := req.DevMode()
	r := c.R()
	resp, err := r.EnableTrace(true).
		SetCookies().
		Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	if resp.IsSuccess() {
		var a TResp
		b := resp.Bytes()
		json.Unmarshal(b, &a)
		//fmt.Printf("%+v", s)
		fmt.Println(s)
	}

	if resp.IsError() {
		resp.String()
	}

	// 特殊处理
	if resp.StatusCode == 123 {

	}
}
