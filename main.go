package main

import (
	"bytes"
	"fmt"
	"forward_openai/app/utils/u_config"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {

	// 代理对象
	forwardProxy := createForwardProxy(u_config.LoadConfig().TargetHost)

	// 这里配置需要转发的path路径
	for i, value := range u_config.LoadConfig().ForwardPathList {
		fmt.Printf("forwardProxy %d: %s\n", i, value)
		http.Handle(value, forwardProxy)
	}

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// 创建转发代理
func createForwardProxy(targetHost string) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{

		// 修改请求体
		Director: func(req *http.Request) {
			// 动态替换请求域名为目标域名地址，path路径自拼接
			oldRqeUrl := fmt.Sprintf("%v%v", req.Host, req.RequestURI)
			targetUrl := fmt.Sprintf("%v%v", targetHost, req.RequestURI)
			log.Printf("【forward】%v => %v\n", oldRqeUrl, targetUrl)
			u, _ := url.Parse(targetUrl)
			req.URL = u
			req.Host = u.Host // 必须显示修改Host，否则转发可能失败

			// 处理请求参数
			parseReqParams(req)
		},

		// 修改响应体
		ModifyResponse: func(resp *http.Response) error {
			/*
				// 这里可以修改响应内容
				log.Println("resp status:", resp.Status)
				log.Println("resp headers:")
				for hk, hv := range resp.Header {
					log.Println(hk, ":", strings.Join(hv, ","))
				}
			*/
			return nil
		},

		// 错误处理
		ErrorLog: log.New(os.Stdout, "ReverseProxy:", log.LstdFlags|log.Lshortfile),
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			if err != nil {
				log.Println("ErrorHandler catch err:", err)

				w.WriteHeader(http.StatusBadGateway)
				_, _ = fmt.Fprintf(w, err.Error())
			}
		},
	}
}

// isJsonReq 是否json格式请求
func isJsonReq(contentType string) bool {
	return len(contentType) > 0 && strings.Contains(contentType, "application/json")
}

// parseReqParams 处理请求参数
func parseReqParams(req *http.Request) {

	if !u_config.LoadConfig().PrintParam {
		// 如果配置了不打印，则跳过
		return
	}

	if isJsonReq(req.Header.Get("Content-Type")) {

		// 打印其他格式请求参数
		log.Println("\nRequest JSON data:")

		data, err := io.ReadAll(req.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
		} else {
			log.Printf("%s\n", data)
		}

		err = req.Body.Close()
		if err != nil {
			log.Println("Error close request body:", err)
		}

		req.Body = io.NopCloser(bytes.NewReader(data))

	} else {

		// 打印其他格式请求参数
		log.Println("\nRequest parameters:")

		if req.Method == "GET" {
			for k, v := range req.Form {
				log.Printf("%v: %v\n", k, v)
			}
		} else if req.Method == "POST" {
			for k, v := range req.PostForm {
				log.Printf("%v: %v\n", k, v)
			}
		}
	}
}
