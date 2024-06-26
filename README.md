## 转发OpenAI接口地址（包含stream流）的小工具
go+docker转发OpenAI的API接口，部署在海外服务器即可对[https://api.openai.com](https://api.openai.com)进行代理，支持stream流。

可以配置（[点击配置文件路径](app/utils/u_config/u_config.go)）：
- 是否打印请求参数，默认值：false
- 转发目标地址，默认值：https://api.openai.com
- 需要转发的路径列表（多个可用|分割），默认值：/v1/chat/completions

### Docker运行
[点击查看使用docker方式运行](docker/README.md)
```docker
docker run -d \
--name forward_openai \
-p 8080:8080 \
-e sg.forward_openai.printParam=false \
-e sg.forward_openai.targetHost="https://api.openai.com" \
-e sg.forward_openai.forwardPathList="/v1/chat/completions" \
--restart=always \
--pull=always \
samge/forward_openai:latest
```

### 使用
将原来的`https://api.openai.com`改为`http://yourIpAddr:8080`即可


### 【可选】源码运行：
```shell
go run main.go
```

### 技术交流
- [Join Discord >>](https://discord.com/invite/eRuSqve8CE)
- WeChat：`SamgeApp`
