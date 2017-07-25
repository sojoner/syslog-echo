FROM golang

WORKDIR /go/src/github.com/sojoner/syslog-echo/

COPY main.go ./

RUN go get \ 
	&& go build

ENV	SYSLOG_HOST=0.0.0.0 \
	SYSLOG_PORT=5140

CMD ["/go/src/github.com/sojoner/syslog-echo/syslog-echo"]

