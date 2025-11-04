FROM golang:1.25-alpine AS build
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

FROM build AS prod
COPY . .
RUN go build -o /usr/local/bin/app ./cmd/kennen/main.go
CMD ["/usr/local/bin/app"]
EXPOSE 8080/tcp