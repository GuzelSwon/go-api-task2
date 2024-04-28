FROM golang:1.22.2 as base_image

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY .. ./

RUN go build -o /app-guzel

EXPOSE 8000

CMD ["/app-guzel"]