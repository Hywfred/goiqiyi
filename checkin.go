package goiqiyi

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Query gets current status of the user specified by p00001
// and stores the value into target.
func Query(p00001 string, target interface{}) {
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

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
