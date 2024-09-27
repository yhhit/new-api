package service

import (
	"crypto/tls"
	"net/http"
	"one-api/common"
	"time"
)

var httpClient *http.Client
var impatientHTTPClient *http.Client

func init() {
	if common.RelayTimeout == 0 {
		httpClient = &http.Client{}
	} else {
		httpClient = &http.Client{
			Timeout: time.Duration(common.RelayTimeout) * time.Second,
		}
	}

	impatientHTTPClient = &http.Client{
		Timeout: 5 * time.Second,
	}
}

func GetHttpClient() *http.Client {
	// 创建一个自定义的Transport，根据环境变量设置TLSClientConfig
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	return client
}

func GetImpatientHttpClient() *http.Client {
	return impatientHTTPClient
}
