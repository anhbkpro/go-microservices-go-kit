package main

import (
	"fmt"
	"github.com/anhbkpro/go-microservices-go-kit/src/api/repositories"
	"github.com/anhbkpro/go-microservices-go-kit/src/api/service"
	"github.com/anhbkpro/go-microservices-go-kit/src/api/transports/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	repo := repositories.NewMariaDBAgentsRepository()
	defer repo.Close()

	svc := service.AgentsServiceImpl{}
	svc.Repo = repo

	insertAgentPlayerEndpoint := endpoints.MakeInsertAgentPlayerEndpoint(svc)

	ph := r.Methods(http.MethodPost).Subrouter()
	ph.Handle("/agent-player", httptransport.NewServer(
		insertAgentPlayerEndpoint,
		endpoints.DecodeInsertAgentPlayerRequest,
		endpoints.EncodeInsertAgentPlayerResponse,
	))

	getAgentByIDEndpoint := endpoints.MakeGetAgentByIdRequestEndpoint(svc)

	gh := r.Methods(http.MethodGet).Subrouter()
	gh.Handle("/agent/{id}", httptransport.NewServer(
		getAgentByIDEndpoint,
		endpoints.DecodeGetAgentByIdRequest,
		endpoints.EncodeGetAgentByIdResponse,
	))

	srv := &http.Server{
		Handler:      ch(r),
		Addr:         ":8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("starting server on port 8080")
	log.Fatal(srv.ListenAndServe())
}
