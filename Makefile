#/bin/bash
# This is how we want to name the binary output
OUTPUT=bin/saas_etcd
SRC=./main.go

# These are the values we want to pass for Version and BuildTime
GITTAG=1.0.0
BUILD_TIME=`date +%Y%m%d%H%M%S`



# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${GITTAG} -X main.Build_Time=${BUILD_TIME} -s -w"

local:
	rm -f ./bin/saas_etcd
	go build ${LDFLAGS} -o ${OUTPUT} ${SRC}

linux:
	rm -f ./bin/saas_etcd
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter ${LDFLAGS} -o ${OUTPUT} ${SRC}


clean:
	rm -f ./bin/saas_etcd
	rm -f ./build/saas_etcd_*
