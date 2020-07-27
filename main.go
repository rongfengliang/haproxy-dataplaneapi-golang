package main

import (
	"log"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/rongfengliang/haproxy-golang/client"
)

var (
	commitId int64
)

func main() {
	// 1. first get  transaction version
	client := client.New(httptransport.New("127.0.0.1:5555", "/v2", []string{"http"}), strfmt.Default)
	// make the authenticated request to get all items
	basicAuth := httptransport.BasicAuth("admin", "dalong")
	frontends, err := client.Frontend.GetFrontends(nil, basicAuth)
	if err != nil {
		log.Panicln("some wrong:", err)
	}
	commitId = frontends.ConfigurationVersion
	log.Println(commitId)

	// for test print frontend info
	for _, frontend := range frontends.GetPayload().Data {
		log.Println("get frontend info :", frontend.Name, frontend.Mode, frontend.DefaultBackend)
	}
	// 2. create transaction

}
