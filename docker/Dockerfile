FROM golang:1.10.3 as builder
WORKDIR /go/src/yonghui.cn/blockchain/ytrace
RUN mkdir -p bin
COPY backend backend
COPY common common
COPY chaincode chaincode
COPY database database
COPY vendor vendor
COPY Makefile Makefile
COPY config config
RUN make build

From ubuntu:latest
RUN apt-get update
RUN apt-get install -y ca-certificates

WORKDIR /opt/ytrace/

COPY --from=builder /go/src/yonghui.cn/blockchain/ytrace/bin/ytrace ./
COPY ./config /opt/ytrace/config
COPY ./templates templates

RUN ln -s /opt/ytrace/ytrace /bin/ytrace
ENTRYPOINT ["ytrace"]
