package u_config

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
)

// ForwardOpenaiConfig 配置
type ForwardOpenaiConfig struct {
	PrintParam      bool     `json:"printParam"`      // 是否打印请求参数，默认：false
	TargetHost      string   `json:"targetHost"`      // 转发目标地址
	ForwardPathList []string `json:"forwardPathList"` // 需要转发的路径列表
}

var config *ForwardOpenaiConfig
var once sync.Once

// ForwardPathListSplitTag ForwardPathList配置的分隔符号
var ForwardPathListSplitTag = "|"

// DefaultTargetHost 默认的转发地址
var DefaultTargetHost = "https://api.openai.com"

// DefaultForwardPath 默认的转发路径
var DefaultForwardPath = "/v1/chat/completions"

// LoadConfig 加载配置
func LoadConfig() *ForwardOpenaiConfig {
	once.Do(func() {
		// 给配置赋默认值
		config = &ForwardOpenaiConfig{
			PrintParam:      false,
			TargetHost:      DefaultTargetHost,
			ForwardPathList: []string{DefaultForwardPath},
		}

		// 判断配置文件是否存在，存在直接JSON读取
		_, err := os.Stat("config.json")
		if err == nil {
			f, err := os.Open("config.json")
			if err != nil {
				log.Fatalf("open config err: %v", err)
				return
			}
			defer func(f *os.File) {
				_ = f.Close()
			}(f)
			encoder := json.NewDecoder(f)
			err = encoder.Decode(config)
			if err != nil {
				log.Fatalf("decode config err: %v", err)
				return
			}
		}

		// 有环境变量使用环境变量
		PrintParam := os.Getenv("sg.forward_openai.printParam")
		if PrintParam != "" {
			config.PrintParam = PrintParam == "true"
		}

		TargetHost := os.Getenv("sg.forward_openai.targetHost")
		if TargetHost != "" {
			config.TargetHost = TargetHost
		}

		ForwardPathList := os.Getenv("sg.forward_openai.forwardPathList")
		if ForwardPathList != "" {
			config.ForwardPathList = strings.Split(ForwardPathList, ForwardPathListSplitTag)
		}
	})

	return config
}
