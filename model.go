package goiqiyi

import (
	"log"
	"os"
	"regexp"
)

type VipInfo struct {
	Code    string  `json:"code"`
	Message string  `json:"msg"`
	VipData VipData `json:"data"`
}

type VipData struct {
	Level            string `json:"level"`
	GrowthValue      string `json:"growthvalue"`
	TodayGrowthValue int    `json:"todayGrowthValue"`
	Deadline         string `json:"deadline"`
}

const (
	p00001regex string = "P00001=(.*?);"
)

type Cookie struct {
	P00001 string
}

// ParseCookie parses cookie and stores p00001 value in target
func ParseCookie(cookie string, target *Cookie) {
	if target == nil {
		log.Fatalln("parameter target is nil.")
	}
	re, err := regexp.Compile(p00001regex)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	res := re.FindStringSubmatch(cookie)
	if len(res) > 1 {
		target.P00001 = res[1]
	}
	print(target.P00001)
}
