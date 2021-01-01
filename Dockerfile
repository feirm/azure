FROM golang:1.15.3-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o azure cmd/azure/main.go
CMD ["./azure"]
EXPOSE 3000