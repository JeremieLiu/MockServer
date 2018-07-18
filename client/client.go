package main

import (
	"net/http"
	"bytes"
	"compress/zlib"

	"log"
	"io/ioutil"
	"fmt"
)

const STATUS_GET_RESPONSE_OK 	= 200
const STATUS_GET_RESPONSE_ERROR = 201
const STATUS_POST_REQUEST_OK     	= 202
const STATUS_POST_REQUEST_ERROR  	= 204
const ERROR_FOR_UNKNOWN			= 100

type status struct {
	code int  `Json:code`
	message string `Json:message`
}

type commitMSG struct {
	status status   `Json:status`
	collect bool    `Json:collect`
	commit bool    `Json:commit`
}


func main() {
	//*****	https get	*****///
	//pool := x509.NewCertPool()
	//caCertPath := "ca.crt"
	//
	//caCrt, err := ioutil.ReadFile(caCertPath)
	//if err != nil {
	//	fmt.Println("ReadFile err:", err)
	//	return
	//}
	//pool.AppendCertsFromPEM(caCrt)
	//
	//cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	//if err != nil {
	//	fmt.Println("Loadx509keypair err:", err)
	//	return
	//}
	//
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{
	//		RootCAs:      pool,
	//		Certificates: []tls.Certificate{cliCrt},
	//	},
	//}
	//client := &http.Client{Transport: tr}
	//resp, err := client.Get("https://localhost:8081")
	//if err != nil {
	//	fmt.Println("Get error:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//************************//

	//mux := http.NewServeMux()
	//
	//handler := http.HandlerFunc(httpsGetStatisConfigQuery())
	//mux.Handler("/v3/api/jgstatisc/collect.cfg?devType=0&net=0&devId=abc",handler)


	//?devType=0&net=0&devId=abc  http尾参数
	http.HandleFunc("/",httpsGetStatisConfigQuery)
	err := http.ListenAndServe("http://10.0.0.76:7070/",nil)
	if err == nil {
		log.Fatal(err)
	}
}

func httpsGetStatisConfigQuery(w http.ResponseWriter , r *http.Request){

	//调用Get请求
	resp, _ := http.Get("http://10.0.0.76:7070/v3/api/jgstatisc/collect.cfg?devType=0&net=0&devId=abc")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://10.0.0.76:7070/v3/api/jgstatisc/collect.cfg?devType=0&net=0&devId=abc", nil)
	request.Header.Set("Connection", "keep-alive")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		log.Print("GET RESPONSE OK")
		log.Println(string(body))
	}else if http.StatusText(r.Response.StatusCode)=="201" {
		body, _ := ioutil.ReadAll(response.Body)
		log.Print("STATUS_GET_RESPONSE_ERROR")
		log.Println(string(body))
	}else{
		body, _ := ioutil.ReadAll(response.Body)
		log.Print("ERROR_FOR_UNKNOWN")
		log.Println(string(body))
	}
}

//func LocalStatisticsReport(){
//	resp, err := http.PostForm("http://example.com/form",
//		url.Values{"key":{"test"},})
//}



//func StatisConfigQuery() {
//
//
//
//	http.Handle("/v3/api/jgstatisc/collect.cfg?devType=0&net=0&devId=abc", scqHandler)
//	var str = []byte(`[
//				"status":{"code":0,"message":"ok"},
//				"collect":true,
//				"commit":false
//                    ]`)
//
//	if bs, err := json.Marshal(user); err == nil {
//		//        fmt.Println(string(bs))
//		req := bytes.NewBuffer([]byte(bs))
//		tmp := `{"name":"junneyang", "age": 88}`
//		req = bytes.NewBuffer([]byte(tmp))
//
//		body_type := "application/json;charset=utf-8"
//		resp, _ = http.Post("http://10.67.2.252:8080/test/", body_type, req)
//		body, _ = ioutil.ReadAll(resp.Body)
//		fmt.Println(string(body))
//	} else {
//
//	}
//
//}

func scqHandler(){
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte("{test1}\n"))
	w.Write([]byte("{test2}\n"))
	w.Write([]byte("{test2}\n"))
	w.Close()
}