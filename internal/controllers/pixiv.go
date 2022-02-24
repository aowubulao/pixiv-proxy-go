package controllers

import (
	"com.neoniou.go/pixiv-proxy-go/internal/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const PixivUrl = "https://i.pximg.net/"

var useProxy = false
var proxyStr string

func init() {
	var proxy = utils.GetArg("proxy")
	if proxy != "" {
		useProxy = true
		proxyStr = proxy
		log.Println("Use proxy: " + proxyStr)
	}
}

func PixivImageProxy(resp http.ResponseWriter, req *http.Request) {
	var uri = req.RequestURI
	if !strings.HasPrefix(uri, "/proxy") {
		utils.BuildOk(resp)
		return
	}
	var pixivUri = strings.Replace(uri, "/proxy", "", 1)
	client := createClient()
	httpReq, err := http.NewRequest("GET", PixivUrl+pixivUri, nil)
	if err != nil {
		utils.BuildOk(resp)
		return
	}
	httpReq.Header.Set("referer", "https://www.pixiv.net/")
	response, err := client.Do(httpReq)
	if err != nil {
		utils.BuildOk(resp)
		return
	}
	resp.Header().Set("content-type", response.Header.Get("content-type"))
	_, err = io.Copy(resp, response.Body)
	if err != nil {
		utils.BuildOk(resp)
		return
	}
}

func createClient() http.Client {
	if useProxy {
		proxyUrl, err := url.Parse(proxyStr)
		if err != nil {
			return http.Client{}
		}
		return http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	} else {
		return http.Client{}
	}
}
