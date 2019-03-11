package main

import (
	"os"
    "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"strings"
	"encoding/base64"
)

type SystemDef struct {
	Name string
	Url string
	BasicAuth string
}

func parse_file(env,filename string) []SystemDef {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var ret []SystemDef
	err = json.Unmarshal(byteValue, &ret)
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

var ExternalSystem = map[string]string {
    "ASK-Auskunft" : "https://services.dev.ask.cp.creditreform.de:8445/ask-auskunft/v1/info/systemInfo",
    "ASK-Verzeichnis" : "https://services.dev.ask.cp.creditreform.de:8447/ask-verzeichnis/v1/info/systemInfo",
    "ASK-Inkasse" : "https://services.dev.ask.cp.creditreform.de:8443/ask-inkasso/v1/info/systemInfo",
}

type SystemInfo struct {
    ServiceName string
    ServiceHealth string
	WarVersion string
	CalledServices map[string]SystemInfo `json:"calledServices"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	 return base64.StdEncoding.EncodeToString([]byte(auth))
  }

func show_system_info(definition SystemDef) {
	fmt.Printf("Testing %s\n",definition.Name)
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
			//fmt.Println(string(data))
			var info SystemInfo
			if err := json.Unmarshal(data, &info); err != nil {
				panic(err)
			}
			fmt.Printf("System info   : %s\n",definition.Name)
			fmt.Printf("Service       : %s\n",info.ServiceName)
			fmt.Printf("Health status : %s\n",info.ServiceHealth)
			fmt.Printf("Version       : %s\n",info.WarVersion)
			fmt.Printf("Services      : %d\n", len(info.CalledServices))
			for cs := range info.CalledServices {
				current := info.CalledServices[cs]
				fmt.Printf("=> Sub-Service: %s\n",current.ServiceName)
				fmt.Printf("  => %s\n",current.ServiceHealth)
			}
		}
    }
}

func main() {
	definitions := parse_file("dev.ask","dev.ask.json")
	fmt.Printf("Defs: %d\n",len(definitions))
	for _, current := range definitions {
		show_system_info(current)
	}
    //for k := range ExternalSystem {
        //show_system_info(k, ExternalSystem[k])
    //}
}
