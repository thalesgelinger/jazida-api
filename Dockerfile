FROM golang:1.23 as builder

RUN apt-get update && apt-get install -y gcc libc6-dev

WORKDIR /app

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o bin/jazida cmd/jazida/main.go

FROM scratch

EXPOSE 8080

COPY --from=builder /app/bin/jazida /jazida
COPY --from=builder /app/web /web

ENTRYPOINT ["/jazida"]

