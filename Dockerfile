FROM golang:1.22

WORKDIR /usr/src/app

RUN go mod init Linkchain

COPY . .

CMD ["go","run","main.go"]
