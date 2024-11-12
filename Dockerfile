FROM golang:1.23.2

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY public ./public
RUN go build -o ./server

EXPOSE 3000
CMD ["./server"]
