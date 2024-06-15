FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/jazida cmd/musico-api/main.go

FROM scratch 
EXPOSE 8080
COPY --from=builder /app/bin/jazida /musico
ENTRYPOINT ["/jazida"]

