## docker build --no-cache=true -f dep.dockerfile -t gokit-lorem-tracing-01 .
FROM golang:1.11

WORKDIR ${GOPATH}/src/github.com/generals-space/gokit/07.gokit链路追踪/01.gokit-lorem-single-tracing
COPY . .

## 192.168.0.8:9999是我本地ss代理地址...
## 环境变量需要unset, 否则会影响示例中的服务连接
RUN export http_proxy=http://192.168.0.8:9999 \
    && export https_proxy=http://192.168.0.8:9999 \
    && go get -v ./server \
    && go get -v ./client \
    && unset http_proxy \
    && unset https_proxy
