package main

import (
	"fmt"
	"os"
	"time"
    "io/ioutil"
    "net/http"
)

func main() {
	f, err := os.OpenFile("egal.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
    	panic(err)
	}

	defer f.Close()
	for i := 0; i < 10; i++ {
		start := time.Now()
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://localhost:8000/supplierinquiries", nil)	
		req.Header.Set("MemberId", "615000191")
		res, err := client.Do(req)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(res.Body)
			f.Write(data)
		}
		elapsed := time.Since(start)
		fmt.Printf("Request took %s\n", elapsed)
	}
}