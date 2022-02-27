package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	//设置version
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)

	//将request中的header设置到reponse中
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}

	// 记录日志并输出
	clientIP := getCurrentIP(r)
	log.Printf("Success! client ip: %s\n", clientIP)
	log.Printf("Success! client response code: %d", 200)
}

func getCurrentIP(r *http.Request) string {
	//ip := r.Header.Get("X-Forwarded-IP")
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		//当请求头不存在即不存在代理时直接获取ip
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("start httpd failed, err: %s\n", err.Error())
	}
}
