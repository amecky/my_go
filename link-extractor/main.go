package main

import (
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
)

// https://services.abcp.creditreform.de:30028/notification-system-admin-gui
func ExtractName(url string) string {
	lastBin := strings.LastIndex( url, ".de" )
	tmp := url[lastBin:len(url)]
	tt := strings.Index(tmp,"/")
	next := tmp[tt + 1:len(tmp)]
	ll := strings.Index(next,"/")
	if ll == -1 {
		return next
	}
	return next[0:ll]
}

func main() {
	b, err := ioutil.ReadFile("ASK+Deployment.txt")
    if err != nil {
        fmt.Print(err)
	}
	var lst []string
	re := regexp.MustCompile("https://([A-Za-z0-9\\.\\-/:]+)")
	res := re.FindAll(b,-1)
	for _,r := range res {
		found := false
		str := string(r)
		for _,current := range lst {
			if current == str {
				found = true
			}
		}
		if !found {			
			if strings.Contains(str,"abcp") {
				lst = append(lst,str)
			}
		}
	}
	for _, c := range lst {
		fmt.Println("{")
		fmt.Printf("\t\"name\" : \"%s\",\n",ExtractName(c))
		fmt.Printf("\t\"url\" : \"%s\",\n",c)
		fmt.Println("},")
	}
	//fmt.Println(content)
}