package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"math/rand"
	"time"
	"encoding/hex"
	"bytes"
	"compress/zlib"
)

const STATUS_GET_RESPONSE_OK 	= 200
const STATUS_GET_RESPONSE_ERROR = 201
const STATUS_POST_REQUEST_OK    = 202
const STATUS_POST_REQUEST_ERROR = 204
const ERROR_FOR_UNKNOWN			= 100


//Json格式设计的内部类
type status struct {
	Code int  `json:Code`
	Message string `json:Message`
}

//定义Json格式类
type commitMSG struct {
	Status status   `json:"Status"`
	Collect bool    `json:Collect`
	Commit bool     `json:Commit`
}

//定义本地统计上传数据格式
type LocalStatisticsReportMSG struct{
	Head string			`json:"Head"`
	Content string		`json:"Content"`
	ContentId string	`json:"Content_id"`
	Rear string			`json:"Rear"`
}



func main(){
	fmt.Println(1)

	// 测试统计配置查询
	// http.HandleFunc("/v3/api/jgstatisc/collect.cfg",httpsRespStatisConfigQuery)

	// 测试本地统计上报
	//http.HandleFunc("/v3/api/jgstatisc/collect.do",LocalStatisticsReport)

	// 测试本地统计获取
	err := http.ListenAndServe("10.0.0.76:7070",nil)
	if err == nil {
		log.Print(err)
	}
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
func httpsRespStatisConfigQuery(w http.ResponseWriter , r *http.Request){

	//测试路径请求发送路径
	//fmt.Println(r.RequestURI)

	//获取随机commit数值
	commitRst :=  getCommitResult()

	//对Json数据赋值
	var msg commitMSG
	msg.Status.Code = 0
	msg.Status.Message = "OK"
	msg.Collect = true
	msg.Commit = commitRst

	//将go类结构体转化为json
	bs , err := json.Marshal(msg)
	if err == nil {
		var respMSG commitMSG
		if err := json.Unmarshal(bs, &respMSG); err == nil {
			log.Print("respMSG:")
			log.Println(respMSG)
			log.Print("resp:")
			//log.Println(resp)
		} else {
			fmt.Println(err)
		}

		w.WriteHeader(STATUS_GET_RESPONSE_OK)
		w.Write(bs)			//没有显示信息原因:没有吧数据缓存写回 w .
	}else{
		fmt.Print("Json Marshal err:")
		fmt.Println(err)
	}
}


/*
 * 本地统计上报
 */
func LocalStatisticsReport(w http.ResponseWriter , r *http.Request){

	//生成随机字符
	singleString := GetOneRandomString()
	fmt.Println("singleString:"+singleString)

	//构造输出数据项
	var lsrMSG LocalStatisticsReportMSG
	lsrMSG.Head = "{"
	lsrMSG.Content = "测试：用户数据项:"
	lsrMSG.ContentId = singleString
	lsrMSG.Rear = "}"

	//将go类结构体转化为json
	lsrByte , errJson := json.Marshal(lsrMSG)
	if errJson == nil {
		var lsrMSG LocalStatisticsReportMSG
		if err := json.Unmarshal(lsrByte, &lsrMSG); err == nil {
			fmt.Print("lsrMSG:")
			fmt.Println(lsrMSG)
			//log.Print("resp:")
			//log.Println(resp)
		} else {
			log.Println(err)
		}
	}else{
		fmt.Println(errJson)
	}

	//压缩格式
	var b bytes.Buffer
	writerTemp := zlib.NewWriter(&b)
	writerTemp.Write([]byte(lsrByte))
	writerTemp.Close()

	w.WriteHeader(STATUS_POST_REQUEST_OK) //返回状态码
	w.Write(b.Bytes())
	fmt.Println(b.Bytes())


}

/*
 * 本地统计获取
 */

func LocalStatisicsacquisition(w http.ResponseWriter, r *http.Request){

	oneKRandomString := GetOneKRandomString()

	w.WriteHeader(STATUS_GET_RESPONSE_OK)
	w.Write([]byte(oneKRandomString))

}

/*
 * 随机生成1个字符
 */
func  GetOneRandomString() string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 1; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/*
 * 随机生成1024个字符
 */
func  GetOneKRandomString() string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 1024; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	res := hex.EncodeToString([]byte(result))

	return res
}


