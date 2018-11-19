BINARY=bin/ytrace
DOCKER_BIN=yonghui/ytrace
GOPACKAGE=yonghui.cn/blockchain/ytrace/backend
CC_PACKAGE=src/yonghui.cn/blockchain/ytrace/
CC_TAR=bin/chaincode.tar.gz
PWD=`pwd`

default:build

build:
	go build -o ${BINARY} ${GOPACKAGE}

docker:
	docker build --rm=true --tag=${DOCKER_BIN} -f ./docker/Dockerfile .

chaincode:
	tar --transform 's,^,${CC_PACKAGE},' -czf  ${CC_TAR} chaincode

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean build docker chaincode