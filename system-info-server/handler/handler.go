package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"system-info-server/domain"
	"strings"
	"encoding/base64"
	"io/ioutil"
	"github.com/gorilla/mux"
	"html/template"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func load_infos(definitions []domain.SystemDef) []domain.InfoData {
	var ret []domain.InfoData
	for _, definition := range definitions {
		client := &http.Client{}
		req, err := http.NewRequest("GET", definition.Url, nil)
		if definition.BasicAuth != "" {
			params := strings.Split(definition.BasicAuth,",")
			req.Header.Add("Authorization","Basic " + basicAuth(params[0],params[1]))
		}
		response, err := client.Do(req)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, ler := ioutil.ReadAll(response.Body)
			if ler != nil {
				fmt.Printf("The HTTP request failed with error %s\n", ler)
			} else {
				var info domain.InfoData
				if  definition.MediaType == "xml" {
					//fmt.Println(string(data))
					if err := xml.Unmarshal(data, &info); err != nil {
						fmt.Println(err)
					} else {
						ret = append(ret,info)
					}
				} else {
					if err := json.Unmarshal(data, &info); err != nil {
						fmt.Println(err)
					} else {
						ret = append(ret,info)
					}
				}
				
			}
		}
	}
	return ret
}

func MakeGetSystemInfos(repository *domain.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		params := mux.Vars(r)
		env := params["env"]
		definitions := repository.Environments[env]
		infos := load_infos(definitions)
		t, err := template.ParseFiles("list.html")
		if err != nil {
			panic(err)
		} else {
			t.Execute(w, infos)
		}
	}
}