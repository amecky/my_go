package main

import (
	"fmt"
	"log"
	"net/http"
	"system-info-server/domain"
	"system-info-server/handler"
	"github.com/gorilla/mux"
)

func main() {

	repo := domain.Repository{}
	repo.Environments = make(map[string][]domain.SystemDef)	
	ret := domain.Load_Environment("dev.ask","dev.ask.json",&repo)
	if  !ret {
		fmt.Println("Error loading environment")
	}
	ret = domain.Load_Environment("cpint","cpint.json",&repo)
	if  !ret {
		fmt.Println("Error loading environment")
	}
	ret = domain.Load_Environment("abcp","abcp.json",&repo)
	if  !ret {
		fmt.Println("Error loading environment")
	}

	router := mux.NewRouter()
	router.HandleFunc("/info/{env}", handler.MakeGetSystemInfos(&repo)).Methods("GET")
	router.HandleFunc("/info/{env}/requests/{service}", handler.MakeGetRequests(&repo)).Methods("GET")
	router.HandleFunc("/info/{env}/requests/{service}/charts", handler.MakeGetCharts(&repo)).Methods("GET")
	router.HandleFunc("/info/{env}/metrics/{service}", handler.MakeGetMetrics(&repo)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}