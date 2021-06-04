
FROM golang:alpine as base

WORKDIR /app/go/base

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get -u github.com/cosmtrek/air
COPY . .

FROM golang:alpine as builder

WORKDIR /app/go/builder

COPY --from=base /app/go/base /app/go/builder
RUN CGO_ENABLED=0 go build main.go

FROM alpine as prod

WORKDIR /app/go/src

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/go/builder/main /app/go/src/main

EXPOSE 8000
CMD [ "/app/go/src/main" ]