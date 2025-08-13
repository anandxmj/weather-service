FROM golang:1.24-alpine AS builder
RUN apk --no-cache add ca-certificates
WORKDIR /weather-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o weather-service ./cmd/server/main.go


FROM gcr.io/distroless/static:nonroot
COPY --from=builder /weather-service/weather-service /weather-service
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER nonroot:nonroot
EXPOSE 3000
ENTRYPOINT ["/weather-service"]
