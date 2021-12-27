package json

import (
	"bufio"
	"github.com/angel-one/go-utils/log"
	"math/rand"
	"os"
	"time"
)

var scripList []string
var prefix = "{\"a\": 1,\"p\": {\"m\": 1,\"e\": 1,\"t\": [ "
var postfix = " ]}}"

func GetJsonStringSubscription() string{
	if scripList == nil || len(scripList) == 0 {
		readScrips()
	}
	val := ""
	for i:=0; i < 5; i++ {
		val = val + "\"" + getScripValue() + "\","
	}
	sz := len(val)
	if sz > 0 && val[sz-1] == ',' {
		val = val[:sz-1]
	}
	return prefix + val + postfix
}

func getScripValue() string {
	l := len(scripList)
	rand.Seed(time.Now().UnixNano())
	n := 1 +rand.Intn(l)
	return scripList[n-1]
}

func readScrips() {
	file, err := os.Open("resources/scrips.txt")
	if err != nil {
		log.Error(nil).Err(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		if val != "" {
			scripList = append(scripList, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error(nil).Err(err)
	}
}
