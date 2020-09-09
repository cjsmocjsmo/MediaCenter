FROM golang:latest AS builder
COPY tvgobs .
RUN export GOPATH=/go/src/tvgobs
RUN go get -v ./src/tvgobs
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/tvgobs

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
RUN \
  mkdir ./static && \
  mkdir ./static/images && \
  mkdir ./static/images/thumbnails
COPY --from=builder /go/main .
RUN \
  mkdir ./fsData && \
  chmod -R +rwx ./fsData && \
  mkdir ./logs && \
  chmod -R +rwx ./logs && \
  echo "Creating log file" > ./logs/tvgobs_log.txt && \
  chmod -R +rwx ./logs/tvgobs_log.txt
CMD ["./main"]
