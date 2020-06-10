package main

import (
	"fmt"
	"net/http"
	"webgo/webserver"
)

func main() {
	//myHandler := MyHandler{}
	//
	////创建Server结构，并详细配置里面的字段
	//server := http.Server{
	//	Addr:              ":8080",
	//	Handler:           &myHandler,
	//	ReadHeaderTimeout: 2 * time.Second,
	//}
	//
	//server.ListenAndServe()
	webserver.WebServerBase()
}

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过详细配置服务器的信息来处理请求！")
}
