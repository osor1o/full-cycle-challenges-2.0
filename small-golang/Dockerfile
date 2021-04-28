FROM golang:alpine as build

RUN mkdir -p /app

WORKDIR /app

COPY go.mod .

COPY hello.go .

RUN go build ./hello.go

FROM scratch

COPY --from=build /app/hello /app/hello

ENTRYPOINT [ "/app/hello" ]
