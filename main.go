package main

import (
	"fmt"
	"net/http"
	"github.com/meng-qiang-vip/gofirst/hello"
)

// 定义一个处理函数，用于处理 HTTP 请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 向客户端返回 "Hello 孟强"
	fmt.Fprintln(w, "Hello CICD 代码已经使用docker-compose代替")
}

func main() {
	hello.GetSay("孟强111")
	hello.Add(10, 30)
	// 注册路由：当访问根路径 "/" 时，调用 helloHandler 函数
	http.HandleFunc("/", helloHandler)

	// 启动 HTTP 服务，监听 8080 端口
	fmt.Println("HTTP 服务已启动，监听地址：http://localhost:8054")
	err := http.ListenAndServe(":8054", nil)
	if err != nil {
		fmt.Printf("启动 HTTP 服务失败: %v", err)
	}
}
