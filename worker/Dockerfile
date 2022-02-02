FROM golang:alpine3.15
WORKDIR /app
RUN mkdir -p ./archive
COPY template/* ./template/*
COPY server.go ./server.go
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download
RUN apk add texlive
RUN go build -o /worker
EXPOSE 8080
CMD ["/worker"]
