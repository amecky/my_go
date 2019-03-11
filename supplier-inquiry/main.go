package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"supplier-inquiry/domain"
	"supplier-inquiry/handler"
)

func main() {

	repo := domain.NewMockRepository()
	router := mux.NewRouter()
	router.HandleFunc("/supplierinquiries", handler.MakeGetSupplierInquiriesHandler(repo)).Methods("GET")
	router.HandleFunc("/supplierinquiries/{id}", handler.MakeGetSupplierInquiryHandler(repo)).Methods("GET")
	router.HandleFunc("/supplierinquiries/{id}/ignore", handler.MakeIgnoreSupplierInquiryHandler(repo)).Methods("POST")
	router.HandleFunc("/supplierinquiries/{id}/reply", handler.MakeReplySupplierInquiryHandler(repo)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

