## docker build --no-cache=true -f dep.dockerfile -t gokit-lorem-ratelimit-02 .
FROM golang:1.11

WORKDIR ${GOPATH}/src/github.com/generals-space/gokit/06.gokit-playground-example/32.gokit-lorem-ratelimit-juju
COPY . .

## 192.168.0.8:7777是我本地ss代理地址...
## 环境变量需要unset, 否则会影响示例中的服务连接
RUN export http_proxy=http://192.168.0.8:7777 \
    && export https_proxy=http://192.168.0.8:7777 \
    && go get -v ./cmd \
    && unset http_proxy \
    && unset https_proxy
