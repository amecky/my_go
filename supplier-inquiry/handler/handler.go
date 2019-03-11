package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"supplier-inquiry/domain"
)

func MakeGetSupplierInquiriesHandler(repository domain.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		mid, err := strconv.Atoi(r.Header.Get("MemberId"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		lst := repository.FindByMemberId(mid)
		if json, err := json.Marshal(lst); err == nil {
			fmt.Fprintf(w, string(json))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

//
// Handler um eine einzelne Lierue für ein Mitglied zu liefern
//
func MakeGetSupplierInquiryHandler(repository domain.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		mid, _ := strconv.Atoi(r.Header.Get("MemberId"))
		inquiry, err := repository.FindByMemberIdAndId(mid, params["id"])
		if err {
			json.NewEncoder(w).Encode(inquiry)
			return
		}
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

//
// Handler zum ignorieren einer Lierue
//
func MakeIgnoreSupplierInquiryHandler(repository domain.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		mid, _ := strconv.Atoi(r.Header.Get("MemberId"))
		repository.SetStatus(mid, params["id"], "IGNORED")
		w.WriteHeader(http.StatusOK)
	}
}

//
// Handler zum beantworten einer Lierue
//
func MakeReplySupplierInquiryHandler(repository domain.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		mid, _ := strconv.Atoi(r.Header.Get("MemberId"))
		repository.SetStatus(mid, params["id"], "REPLIED")
		w.WriteHeader(http.StatusOK)
	}
}
