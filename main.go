package main

import (
	"context"
	"log"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/rongfengliang/haproxy-golang/client"
	"github.com/rongfengliang/haproxy-golang/client/backend"
	"github.com/rongfengliang/haproxy-golang/client/bind"
	"github.com/rongfengliang/haproxy-golang/client/frontend"
	"github.com/rongfengliang/haproxy-golang/client/server"
	"github.com/rongfengliang/haproxy-golang/client/transactions"
	"github.com/rongfengliang/haproxy-golang/models"
)

// String point
func String(s string) *string {
	return &s
}

// Bool point
func Bool(s bool) *bool {
	return &s
}

// Int64 point
func Int64(s int64) *int64 {
	return &s
}

func main() {

	// context object for  control connet
	ctx := context.Background()

	// 1. first get  transaction version
	client := client.New(httptransport.New("localhost:5555", "/v2", []string{"http"}), strfmt.Default)
	// make the authenticated request to get all items
	basicAuth := httptransport.BasicAuth("admin", "dalong")
	frontends, err := client.Frontend.GetFrontends(nil, basicAuth)
	if err != nil {
		log.Panicln("some wrong:", err)
	}
	versionID := frontends.ConfigurationVersion
	log.Println(versionID)

	// for test print frontend info
	for _, frontend := range frontends.GetPayload().Data {
		log.Println("get frontend info :", frontend.Name, frontend.Mode, frontend.DefaultBackend)
	}
	// 2. create transaction for generate id

	commitTransaction, err := client.Transactions.StartTransaction(&transactions.StartTransactionParams{Version: versionID, Context: ctx}, basicAuth)
	if err != nil {
		log.Panicln("create transaction some wrong:", err.Error())
	}
	transactionID := commitTransaction.Payload.ID
	log.Println(transactionID)

	// 3. create backend object
	createBackendCreated, createBackendAccepted, err := client.Backend.CreateBackend(&backend.CreateBackendParams{
		Context: ctx,
		Data: &models.Backend{
			Name: "dalongrong",
			Mode: "http",
			Balance: &models.Balance{
				Algorithm: String("roundrobin"),
			},
			Httpchk: &models.Httpchk{
				Method:  "HEAD",
				URI:     "/",
				Version: "HTTP/1.1",
			},
		},
		TransactionID: &transactionID,
	}, basicAuth)
	if err != nil {
		log.Panicln("create backend some wrong:", err.Error())
	}
	if createBackendCreated == nil {
		log.Println(createBackendAccepted.GetPayload())
	}
	// 4. add backend server to backend use create server wrapper
	createServerCreated, createServerAccepted, err := client.Server.CreateServer(&server.CreateServerParams{
		TransactionID: &transactionID,
		Backend:       "dalongrong",
		Context:       ctx,
		Data: &models.Server{
			Name:    "server1",
			Address: "10.15.0.80",
			Port:    Int64(80),
			Check:   "enabled",
			Maxconn: Int64(30),
			Weight:  Int64(100),
		},
	}, basicAuth)
	if err != nil {
		log.Panicln("add  backend server has some wrong:", err.Error())
	}
	if createServerCreated == nil {
		log.Println(createServerAccepted.GetPayload())
	}

	// 5. create frontend
	createFrontendCreated, createFrontendAccepted, err := client.Frontend.CreateFrontend(&frontend.CreateFrontendParams{

		TransactionID: &transactionID,
		Data: &models.Frontend{
			Name:           "dalongrong_frontend",
			Mode:           "http",
			DefaultBackend: "dalongrong",
			Maxconn:        Int64(1000),
		},
		Context: ctx,
	}, basicAuth)
	if err != nil {
		log.Panicln("create  frontend  has some wrong:", err.Error())
	}
	if createFrontendCreated == nil {
		log.Println(createFrontendAccepted.GetPayload())
	}

	// 6. create frontend && backend bind
	createBindCreated, createBindAccepted, err := client.Bind.CreateBind(&bind.CreateBindParams{
		TransactionID: &transactionID,
		Frontend:      "dalongrong_frontend",
		Data: &models.Bind{
			Name:    "httpdemo",
			Address: "*",
			Port:    Int64(9000),
		},
		Context: ctx,
	}, basicAuth)
	if err != nil {
		log.Panicln("create  bind  has some wrong:", err.Error())
	}
	if createBindCreated == nil {
		log.Println(createBindAccepted.GetPayload())
	}

	// 7. apply transaction && do commit
	commitTransactionOK, commitTransactionAccepted, err := client.Transactions.CommitTransaction(&transactions.CommitTransactionParams{
		ID:      transactionID,
		Context: ctx,
	}, basicAuth)
	if err != nil {
		log.Panicln("apply haproxy config   has some wrong:", err.Error())
	}
	if commitTransactionOK != nil {
		// apply config ok
		log.Println(commitTransactionOK.GetPayload())
	}
	if commitTransactionAccepted != nil {
		// accept but not apply success
		log.Println(commitTransactionAccepted.GetPayload())
	}

}
