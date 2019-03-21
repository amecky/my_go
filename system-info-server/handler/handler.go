package handler

import (	
	"net/http"
	"system-info-server/domain"
	"system-info-server/view"	
	"github.com/gorilla/mux"
	"html/template"
	"strconv"
)

func MakeGetSystemInfos(repository *domain.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		params := mux.Vars(r)
		env := params["env"]
		definitions := repository.Environments[env]
		var view view.ListView
		view.Environment = env
		view.Infos = LoadInfos(definitions, env)		
		t := template.Must(template.New("list.html").Funcs(template.FuncMap{
			"mod": func(i, j int) bool { return i + 1 % j == 0 },
		}).ParseFiles("list.html"))
		t.Execute(w, view)
	}
}

func MakeGetRequests(repository *domain.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		params := mux.Vars(r)
		env := params["env"]
		service := params["service"]
		var view view.RequestsView
		view.Service = service
		view.Environment = env
		definitions := repository.Environments[env]
		for _,current := range definitions {
			if current.Name == service {
				view.Requests = LoadRequests(current)
			}
		}
		t := template.Must(template.New("requests.html").Funcs(template.FuncMap{
			"mod": func(i, j int) bool { return i + 1 % j == 0 },
		}).ParseFiles("requests.html"))
		t.Execute(w, view)
	}
}

func MakeGetMetrics(repository *domain.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		params := mux.Vars(r)
		env := params["env"]
		service := params["service"]
		var view view.MetricsView
		view.Service = service
		view.Environment = env
		definitions := repository.Environments[env]
		for _,current := range definitions {
			if current.Name == service {
				view.Metrics = LoadMetrics(current)
			}
		}
		t := template.Must(template.New("metrics.html").ParseFiles("metrics.html"))
		t.Execute(w, view)
	}
}

func MakeGetCharts(repository *domain.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		params := mux.Vars(r)
		env := params["env"]
		service := params["service"]
		path,_ := r.URL.Query()["path"]
		var cv view.ChartsView
		cv.Path = path[0]
		cv.Environment = env
		definitions := repository.Environments[env]
		for _,current := range definitions {
			if current.Name == service {
				requests := LoadRequests(current)
				for idx, r := range requests {
					if path[0] == r.Path {
						if idx < 25 {
							var pd view.PerformanceData
							pd.Key = strconv.Itoa(idx)
							pd.Value = r.Elapsed
							pd.RequestTimestamp = r.RequestTimestamp
							cv.Data = append(cv.Data,pd)
						}
					}
				}
			}
		}
		t := template.Must(template.New("charts.html").ParseFiles("charts.html"))
		t.Execute(w, cv)
	}
}