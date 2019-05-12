## docker build --no-cache=true -f dep.dockerfile -t gokit-lorem-restful-02 .
FROM golang:1.11

WORKDIR ${GOPATH}/src/github.com/generals-space/gokit/06.gokit-playground-example/02.gokit-lorem-restful-ServerErrorEncoder
COPY . .

## 192.168.0.8:7777是我本地ss代理地址...
## 环境变量需要unset, 否则会影响示例中的服务连接
RUN export http_proxy=http://192.168.0.8:7777 \
    && export https_proxy=http://192.168.0.8:7777 \
    && go get -v ./cmd \
    && unset http_proxy \
    && unset https_proxy
