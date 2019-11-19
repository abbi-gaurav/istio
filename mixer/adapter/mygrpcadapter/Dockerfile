FROM golang:1.13 as builder

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo -v -o bin/mygrpcadapter ./cmd/

FROM alpine:3.8
RUN apk --no-cache add ca-certificates

WORKDIR /bin/

COPY --from=builder /app/bin/mygrpcadapter .
ENTRYPOINT [ "/bin/mygrpcadapter" ]
CMD [ "44225" ]
EXPOSE 44225