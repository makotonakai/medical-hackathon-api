FROM golang:1.20

WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

EXPOSE 1323
CMD ["/main"]
