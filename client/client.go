package main
import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
	"compress/zlib"
)

func main() {
	pool := x509.NewCertPool()
	caCertPath := "ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

//func LocalStatisticsReport(){
//	resp, err := http.PostForm("http://example.com/form",
//		url.Values{"key":{"test"},})
//}

func StatisConfigQuery(){
	http.Handle("/v3/api/jgstatisc/collect.cfg",scqHandler)

}

func scqHandler(){
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte("{test1}\n"))
	w.Write([]byte("{test2}\n"))
	w.Write([]byte("{test2}\n"))
	w.Close()
}