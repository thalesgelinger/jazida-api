FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o bin/jazida cmd/jazida/main.go

FROM scratch 
EXPOSE 8080
COPY --from=builder /app/bin/jazida /jazida
COPY --from=builder /app/templates /templates
ENTRYPOINT ["/jazida"]

