test:
	go clean -testcache
	go test -v ./...

tidy:
	go mod tidy