FROM golang:1.21 as build

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o fetchweb

RUN chmod +x fetchweb

CMD [ "./fetchweb" ]