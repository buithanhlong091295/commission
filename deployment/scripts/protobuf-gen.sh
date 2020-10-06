#!/bin/bash
DIR=$GOPATH/src/xtek/exchange/commission
SERVICE_NAME=commission
TEMP_DIR=$DIR/.proto_temp

cp -R $DIR/protobuf $TEMP_DIR

find $TEMP_DIR -name "*.proto" -exec sed -i.bak "s/__SERVICE_NAME__/$SERVICE_NAME/" {} +

IMPORT="-I$TEMP_DIR/ \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate"

rm -rf $DIR/pb
# rm $DIR/pb/**/*.pb.go
# rm $DIR/pb/**/*.pb.gw.go
# rm $DIR/docs/*.swagger.json
# rm -rf $DIR/docs/*

for pkg in $(ls -c $TEMP_DIR | grep -v google); do
  PROTO=$TEMP_DIR/$pkg/*.proto

  protoc $IMPORT --go_out=plugins=grpc:$GOPATH/src/. --validate_out="lang=go:$GOPATH/src/." $TEMP_DIR/$pkg/dto/*.proto
  protoc $IMPORT --go_out=plugins=grpc:$GOPATH/src/. --validate_out="lang=go:$GOPATH/src/." $TEMP_DIR/$pkg/types/*.proto

  protoc $IMPORT --go_out=plugins=grpc:$GOPATH/src/. --validate_out="lang=go:$GOPATH/src/." $PROTO

done

rm -rf $TEMP_DIR