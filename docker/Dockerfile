FROM golang:alpine as builder

# Cacheable steps
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

# Download deps
RUN go mod download
COPY . .

# Environment
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Test
RUN  go test ./...

# Build
RUN go build -a -o /bin/cadence-metrics epam.com/cadence-metrics/cmd/metrics

FROM alpine
EXPOSE 9010
COPY --from=builder /bin/cadence-metrics /usr/bin/
# default configuration
COPY config.yaml /etc/cadence-metrics/config.yaml
CMD ["/usr/bin/cadence-metrics", "-config", "/etc/cadence-metrics/config.yaml"]
