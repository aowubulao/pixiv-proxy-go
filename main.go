package main

import (
	"com.neoniou.go/pixiv-proxy-go/internal/controllers"
	"com.neoniou.go/pixiv-proxy-go/internal/utils"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.PixivImageProxy)
	var addressWithPort = "0.0.0.0:" + utils.GetArgWithDefault("port", "10901")
	log.Println("Listen to " + addressWithPort)
	err := http.ListenAndServe(addressWithPort, nil)
	if err != nil {
		log.Println("Create server error: " + err.Error())
		return
	}
}
