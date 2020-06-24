package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://dev-services.umarkets.com/srvgtw/autochartist/v1/chartpatterns", nil)

	req.Header.Add("SID", "active session id")
	req.Header.Add("UMID", "current um session")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
