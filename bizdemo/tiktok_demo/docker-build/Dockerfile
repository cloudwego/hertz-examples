FROM golang:1.18-alpine3.17 AS build

COPY biz/ /go/src/biz/
COPY pkg/ /go/src/pkg/
COPY go.mod go.sum *.go /go/src/

WORKDIR "/go/src/"
RUN go env -w GO111MODULE=on \
  && go env -w GOPROXY=https://goproxy.cn,direct \
  && go env -w GOOS=linux \
  && go env -w GOARCH=amd64
RUN go mod tidy
RUN go build -o offer_tiktok


FROM jrottenberg/ffmpeg:5-alpine

RUN mkdir "/app"
COPY --from=build /go/src/offer_tiktok /app/offer_tiktok

RUN chmod +x /app/offer_tiktok

EXPOSE 18005
ENTRYPOINT ["/app/offer_tiktok"]