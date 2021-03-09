FROM golang:alpine as builder

COPY src/main.go .

ENV CGO_ENABLED=0

RUN go build -ldflags="-s -w" main.go

FROM scratch

COPY --from=builder /go/main .

EXPOSE 8000

ENTRYPOINT [ "./main" ]
