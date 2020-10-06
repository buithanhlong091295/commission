# Start from golang v1.11 base image
FROM xtek/golang_grpc:latest as build-stage

WORKDIR $GOPATH/src/xtek/exchange/commission
COPY . .

RUN make proto-gen

RUN make build-alpine

FROM alpine:3.7 as production-stage

COPY --from=build-stage /go/src/xtek/exchange/commission/commission_srv /bin