FROM golang:1.18

RUN     mkdir -p /app
WORKDIR /app
COPY . .

RUN go mod download
RUN go build ./main.go

EXPOSE 8080
CMD ["./main"]