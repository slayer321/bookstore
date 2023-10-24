package server

import (
	"log"
	"net/http"

	"github.com/slayer321/bookstore/api/gen/book/v1/bookv1connect"
)

const address = "localhost:8080"

type bookStoreServiceServer struct {
	bookv1connect.UnimplementedInventoryServiceHandler
}

func CreateServer() {

	mux := http.NewServeMux()
	path, handler := bookv1connect.NewInventoryServiceHandler(&bookStoreServiceServer{})
	mux.Handle(path, handler)

	log.Printf("Listening on %s", address)
	http.ListenAndServe(address, mux)
}
