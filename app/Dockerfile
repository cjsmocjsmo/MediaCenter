FROM arm32v7/golang:1.12.13 AS builder
COPY moviegobs .
RUN export GOPATH=/go/src/moviegobs
RUN go get -v ./src/moviegobs
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/moviegobs

FROM arm32v6/alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/main .
RUN \
  mkdir ./static && \
  chmod -R +rwx ./static && \
  mkdir ./fsData && \
  chmod -R +rwx ./fsData && \
  mkdir ./logs && \
  chmod -R +rwx ./logs && \
  echo "Creating log file" > ./logs/moviegobs_log.txt && \
  chmod -R +rwx ./logs/moviegobs_log.txt
  
CMD ["./main"]
