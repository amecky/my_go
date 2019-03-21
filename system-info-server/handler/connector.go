package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"system-info-server/domain"
	"strings"
	"io/ioutil"
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
)

func LoadMetrics(definition domain.SystemDef) []domain.Metric {
	var ret []domain.Metric
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", definition.Url + "/metrics", nil)
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
			if err := json.Unmarshal(data, &ret); err != nil {
				fmt.Println(err)
			} 		
		}
	}
	return ret
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func LoadInfos(definitions []domain.SystemDef, env string) []domain.InfoData {
	var ret []domain.InfoData
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	for _, definition := range definitions {
		var info domain.InfoData
		info.Url = definition.Url
		info.SupportMetrics = definition.SupportMetrics
		info.Environment = env
		req, err := http.NewRequest("GET", definition.Url + "/systemInfo", nil)
		if definition.BasicAuth != "" {
			params := strings.Split(definition.BasicAuth,",")
			req.Header.Add("Authorization","Basic " + basicAuth(params[0],params[1]))
		}
		response, err := client.Do(req)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			info.ServiceHealth = "Down"
			info.ServiceName = definition.Name
			ret = append(ret,info)
		} else {
			data, ler := ioutil.ReadAll(response.Body)
			if ler != nil {
				fmt.Printf("The HTTP request failed with error %s\n", ler)
				info.ServiceHealth = "Down"
				info.ServiceName = definition.Name
				ret = append(ret,info)
			} else {
				
				if  definition.MediaType == "xml" {
					if err := xml.Unmarshal(data, &info); err != nil {
						fmt.Println(err)
						info.ServiceName = definition.Name
						info.ServiceHealth = "Down"
						ret = append(ret,info)
					} else {
						info.ServiceName = definition.Name
						ret = append(ret,info)
					}
				} else {
					if err := json.Unmarshal(data, &info); err != nil {
						fmt.Println(err)
						info.ServiceHealth = "Down"
						info.ServiceName = definition.Name
						ret = append(ret,info)
					} else {
						info.ServiceName = definition.Name
						ret = append(ret,info)
						
					}
				}
				
			}
		}
	}
	return ret
}

func LoadRequests(definition domain.SystemDef) []domain.CustomLogData {
	var ret []domain.CustomLogData
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", definition.Url + "/requests", nil)
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
			if err := json.Unmarshal(data, &ret); err != nil {
				fmt.Println(err)
			} 		
		}
	}
	return ret
}