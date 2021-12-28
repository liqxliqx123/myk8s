package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header.Clone() {
		for _, v1 := range v {
			w.Header().Add(k, v1)
		}
	}
	//

	//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	err2 := os.Setenv("VERSION", "2022")
	if err2 != nil {
		log.Fatal(err2)
	}
	version := os.Getenv("VERSION")
	w.Header().Add("VERSION", version)

	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出

	fmt.Println(w.Header())
	_, err := w.Write([]byte("Hello " +  version))
	if err != nil {
		log.Fatal(err)
	}

	clientIP := ClientIP(r)
	log.Printf("客户端IP: %s, HTTP返回: %d", clientIP, http.StatusOK)

}

func healthz(w http.ResponseWriter, r *http.Request)(){
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}


func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}