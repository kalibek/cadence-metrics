test:
	go test ./...

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /bin/cadence-metrics epam.com/cadence-metrics/cmd/metrics

docker_build:
	docker build . -f docker/Dockerfile -t cadence-metrics