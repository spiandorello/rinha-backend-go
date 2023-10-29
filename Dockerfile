FROM golang:1.21 as build_deps

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -v -ldflags='-s -w' -o /usr/local/bin/app ./cmd/main.go

FROM build_deps

WORKDIR /usr/src/app

EXPOSE 8085

CMD ["app"]