package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// request post
func httpDo(method string, url string, msg string, token string) string {
	client := &http.Client{}
	body := bytes.NewBuffer([]byte(msg))
	reqest, err := http.NewRequest(method, url, body) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	//Add 头协议
	reqest.Header.Add("Accept", "application/json; charset=UTF-8")
	reqest.Header.Add("Content-Type", "application/json")
	// reqest.Header.Add("ServerVersion", "12")
	if token != "" {
		reqest.Header.Add("Authorization", token)
	}
	response, err := client.Do(reqest) //提交
	defer response.Body.Close()

	rest, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		// handle error
	}
	str := string(rest)
	return str
}

func POST(url string, body string, token string) string {
	res := httpDo("POST", url, body, token)
	// fmt.Println(res)
	return res
}

func GET(url string, body string, token string) string {
	res := httpDo("GET", url, body, token)
	// fmt.Println(res)
	return res
}
func test() {
	for i := 1; i < 10000; i++ {
		url := "http://192.168.1.22:9090/cypher"
		token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJmbWEuYWkiLCJwYXNzd29yZCI6ImFkbWluMTIzNDU2IiwidXNlciI6ImFkbWluIn0.gYqjoJYejyuypdcq0asdWEEwzGnEEpPqzer6W_J-kC4"
		body := "{\"graph\": \"default\",\"script\": \"call dbms.config.update({OPT_DB_ASYNC:false, OPT_TXN_OPTIMISTIC:true, OPT_AUDIT_LOG_ENABLE:true})\"}"
		res := POST(url, body, token)
		println(res)

		url = "http://192.168.1.22:9090/db/default/python_plugin/py_test"
		body = "{\"data\":\"\", \"timeout\":0}"
		res2 := POST(url, body, token)
		println(res2)

		url = "http://192.168.1.22:9090/db/default/import/text"

		body = "{\"description\": \"LABEL=company\\nscale:INT32:ID,address:STRING,name:STRING\",\"data\": \"" + strconv.Itoa(i) + ",清华西路,费马\\n\",\"continue_on_error\": true}"
		println(body)
		res1 := POST(url, body, token)
		println(res1)
		url = "http://192.168.1.22:9090/db/default/node/11"
		res = GET(url, "", token)
		fmt.Println(res)
	}

}

func main() {
	for i := 1; i < 20; i++ {
		go test()
	}
	time.Sleep(time.Duration(500) * time.Second)

}
