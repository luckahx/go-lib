#!/bin/sh

# plugins needed:
# go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
# go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
# go get -u github.com/golang/protobuf/protoc-gen-go

DIR=$(cd $(dirname ${0}) && pwd)
cd $DIR
echo "BASE DIR: $DIR"

for d in *
do
    [ ! -d "$DIR/$d" ] && continue
    echo "Scanning $d for *.proto"
    cd $DIR/$d
    for i in *.proto
    do 
         protoc -I . \
            -I$GOPATH/src \
            -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
            -I$GOPATH/src/github.com/luckahx/go-lib/cerr/proto \
            --go_out=plugins=grpc:. \
            --swagger_out=logtostderr=true:. \
            $i
    done
done
            # -I$GOPATH/src/github.com/hearsmart/go-dbtrack/shared/protobuf \