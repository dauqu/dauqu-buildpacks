
	FROM golang:1.16-alpine
	WORKDIR /Users/harshaweb/Documents/go/vhost
	COPY go.mod ./
	COPY go.sum ./
	COPY *.go ./
	RUN go get
	RUN go build main.go
	EXPOSE 8000
	CMD [ "./main" ]
	