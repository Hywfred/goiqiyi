package goiqiyi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Query gets current status of the user specified by p00001
func Query(p00001 string) {
	req, err := http.NewRequest("GET", "https://serv.vip.iqiyi.com/vipgrowth/query.action", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("P00001", p00001)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(data))

}
