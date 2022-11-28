FROM golang:alpine as base

RUN apk add --no-cache git ca-certificates
WORKDIR /var/fargate-poc
COPY . .
RUN CGO_ENABLED=0 go build -ldflags '-s -w' -o main .

FROM scratch
WORKDIR /root/
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base /var/fargate-poc/main .
CMD ["./main"]