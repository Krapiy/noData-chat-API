FROM golang:1.10.3
RUN mkdir -p /go/src/github.com/Krapiy/noData-chat-API
WORKDIR /go/src/github.com/Krapiy/noData-chat-API
COPY . /go/src/github.com/Krapiy/noData-chat-API
CMD ["make", "run_development"]
