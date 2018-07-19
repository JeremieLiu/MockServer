package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"bytes"
	"compress/zlib"
	"math/rand"
	"time"
)

//定义本地统计上传数据格式
type LocalStatisticsReportMSG struct{
	Head string			`json:"Head"`
	Content string		`json:"Content"`
	ContentId string	`json:"Content_id"`
	Rear string			`json:"Rear"`
}

func main(){
	//uriHttpRespStatisConfigQuery()
	uriLocalStatisticsReport()
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


//统计服务上报访问方法
func uriHttpRespStatisConfigQuery()  {

	client := &http.Client{}

	req, err := http.NewRequest("GET","http://localhost:7070/v3/api/jgstatisc/collect.cfg?devType=0&net=0&devId=abc",nil)

	response, err := client.Do(req)
	defer response.Body.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("HttpRespStatisConfigQuery succeed")
}

func uriLocalStatisticsReport(){


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
	fmt.Print("WriterTemp:")
	fmt.Println(writerTemp)
	//log.Println(writerTemp)
	writerTemp.Close()


	body := bytes.NewReader(b.Bytes())
	fmt.Print("Body:")
	fmt.Println(body)
	client := &http.Client{}
	req, err := http.NewRequest("POST","http://localhost:7070/v3/api/jgstatisc/collect.do?devId=abc", body)
	if err != nil {
		log.Println(err)
		log.Println(req)
	}
	fmt.Println("NewRequest succeed")

	response, err := client.Do(req)
	if response != nil {
		fmt.Println(response)
	}
	defer response.Body.Close()
	fmt.Println("HttpRespStatisConfigQuery succeed")
}