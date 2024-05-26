FROM golang:1.22.3-apline

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /pricfetcher

EXPOSE 3000

CMD ["/pricfetcher"]