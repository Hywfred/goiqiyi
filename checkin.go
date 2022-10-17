package goiqiyi

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
	"unicode/utf16"
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

// GenerateRandomStr spawns a random string
func GenerateRandomStr(len int) string {
	rand.Seed(time.Now().UnixNano())
	str := ""
	for i := 0; i < len; i++ {
		rdm62 := uint16(math.Floor(rand.Float64() * 62))
		var code uint16 = 0
		if rdm62 < 10 {
			code = rdm62 + 48
		} else if rdm62 < 36 {
			code = rdm62 + 55
		} else {
			code = rdm62 + 61
		}
		coderune := utf16.Decode([]uint16{code})[0]
		str += string(coderune)
	}
	return str
}
