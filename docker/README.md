## 使用docker操作forward_openai

### 编译docker基础镜像-
```shell
docker build -t samge/forward_openai:base -f docker/Dockerfile_base .
```

### 编译docker镜像
```shell
docker build -t samge/forward_openai -f docker/Dockerfile .
```

### 使用docker运行
- 使用默认配置快速运行
```shell
docker run -itd \
 --name forward_openai \
 -p 8080:8080 \
 --restart=always \
 --pull=always \
 samge/forward_openai:latest
```


你也可以使用指定配置参数运行本项目：

`第一种：基于环境变量运行`

```sh
### 运行项目，环境变量参考下方配置说明，如果配置多个 forwardPathList ，可以使用|分割
docker run -itd \
 --name forward_openai \
 -p 8080:8080 \
 -e sg.forward_openai.printParam="true" \
 -e sg.forward_openai.targetHost="https://api.openai.com" \
 -e sg.forward_openai.forwardPathList="/v1/chat/completions|/v1/chat/path2|/v1/chat/path3" \
 --restart=always \
 --pull=always \
 samge/forward_openai:latest
```

运行命令中映射的配置文件参考下边的配置文件说明。

`第二种：基于配置文件挂载运行`

```sh
### 复制配置文件，根据自己实际情况，调整配置里的内容
sudo mkdir -p `pwd`/docker_data/forward_openai
cp config.dev.json `pwd`/docker_data/forward_openai/config.json  # 其中 config.dev.json 从项目的根目录获取

### 运行项目
docker run -itd \
 --name forward_openai \
 -v `pwd`/docker_data/forward_openai/config.json:/app/config.json \
 -p 8080:8080 \
 --restart=always \
 --pull=always \
 samge/forward_openai:latest
```
