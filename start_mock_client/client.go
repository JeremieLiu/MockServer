package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

const STATUS_GET_RESPONSE_OK 	= 200
const STATUS_GET_RESPONSE_ERROR = 201
const STATUS_POST_REQUEST_OK    = 202
const STATUS_POST_REQUEST_ERROR = 204
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

func main(){


	http.HandleFunc("/v3/api/jgstatisc/collect.cfg?devType=0&net=0&devId=abc",httpsGetStatisConfigQuery)
	http.HandleFunc("/v3/api/jgstatisc/test",httpShowRequest)

	if err := http.ListenAndServe("http://10.0.0.76:7070", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func httpsGetStatisConfigQuery(w http.ResponseWriter , r *http.Request){
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://10.0.0.76:7070/v3/api/jgstatisc/collect.cfg?devType=0&net=0&devId=abc", nil)
	response, _ := client.Do(request)

	if response.StatusCode == STATUS_GET_RESPONSE_OK {
		body, _ := ioutil.ReadAll(response.Body)
		log.Print("GET RESPONSE OK")
		log.Println(string(body))
	}else if http.StatusText(r.Response.StatusCode) == string(STATUS_GET_RESPONSE_ERROR){
		body, _ := ioutil.ReadAll(response.Body)
		log.Print("STATUS_GET_RESPONSE_ERROR")
		log.Println(string(body))
	}else{
		body, _ := ioutil.ReadAll(response.Body)
		log.Print("ERROR_FOR_UNKNOWN")
		log.Println(string(body))
	}
}

func httpShowRequest(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	body_str := string(body)
	w.Write([]byte("Request"))
	w.Write([]byte(body))
	log.Println(body_str)
}