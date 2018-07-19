package main

import (
    "io/ioutil"
    "net/http"
    "bytes"
    "math/rand"
    "encoding/json"
    "log"
    "fmt"
)

const STATUS_GET_RESPONSE_OK 	= 200
const STATUS_GET_RESPONSE_ERROR = 201
const STATUS_POST_REQUEST_OK     	= 202
const STATUS_POST_REQUEST_ERROR  	= 204
const ERROR_FOR_UNKNOWN			= 100


//Json格式设计的内部类
type status struct {
    code int  `Json:code`
    message string `Json:message`
}

//定义Json格式类
type commitMSG struct {
    status status   `Json:status`
    collect bool    `Json:collect`
    commit bool    `Json:commit`
}

// https 测试函数
//type myhandler struct {
//}
//
//func (h *myhandler) ServeHTTP(w http.ResponseWriter,
//    r *http.Request) {
//    fmt.Fprintf(w,
//        "Hi, This is an example of http service in golang!\n")
//}

func main() {
    //*****    HTTPS POST     *****//
    //pool := x509.NewCertPool()
    //caCertPath := "ca.crt"
	//
    //caCrt, err := ioutil.ReadFile(caCertPath)
    //if err != nil {
    //    fmt.Println("ReadFile err:", err)
    //    return
    //}
    //pool.AppendCertsFromPEM(caCrt)
	//
    //s := &http.Server{
    //    Addr:    ":8081",
    //    Handler: &myhandler{},
    //    TLSConfig: &tls.Config{
    //        ClientCAs:  pool,
    //        ClientAuth: tls.RequireAndVerifyClientCert,
    //    },
    //}
	//
    //err = s.ListenAndServeTLS("server.crt", "server.key")
    //if err != nil {
    //    fmt.Println("ListenAndServeTLS err:", err)
    //}
    //***************************//

    httpsRespStatisConfigQuery()
}

/*
 * 随机数生成commit值
 */
func getCommitResult () bool{
    var res bool
    temp := rand.Intn(100)
    if temp%2 ==1 {
        res = false
    }else {
        res = true
    }
    return res
}

/*
 * 统计配置查询
 */

func httpsRespStatisConfigQuery(){

    //获取随机commit数值
    commitRst :=  getCommitResult()
    //构建变量ResposeWriter 用于编写StatusCode
    var w http.ResponseWriter

    //对Json数据赋值
    var msg commitMSG
    msg.status.code = 0
    msg.status.message = "OK"
    msg.collect = true
    msg.commit = commitRst

    //将go类结构体转化为json
    bs , err := json.Marshal(msg)
    if err == nil {

        //将json转化为二进制字符
        req := bytes.NewBuffer([]byte(bs))
        body_type := "application/json;charset=utf-8"

        resp, _ := http.Post("http://10.0.0.76/v3/api/jgstatisc/collect.cfg" , body_type , req)

        body, _ := ioutil.ReadAll(resp.Body)
        log.Println("post commit log:"+ string (body) )

        w.WriteHeader(STATUS_GET_RESPONSE_OK)

        //关闭post请求
        resp, errGet := http.Get("http://10.0.0.76/v3/api/jgstatisc/collect.cfg")
        if errGet != nil {
            fmt.Print("get url error:")
            fmt.Println(errGet)
        }
        defer resp.Body.Close()

        //body, err := ioutil.ReadAll(resp.Body)
    }else{
        fmt.Print("Json Marshal err:")
        fmt.Println(err)
    }

    //var b bytes.Buffer
    //r, errJson := zlib.NewReader(&b)
    //if errJson == nil {
    //    fmt.Println(errJson)
    //}
    //io.Copy(os.Stdout, r)
    //r.Close()
}

