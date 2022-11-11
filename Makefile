v_test:
	go clean -testcache
	go test -v ./...

consumer:
	go run cmd/consumer/main.go

producer:
	go run cmd/producer/main.go

test:
	go test ./...

tidy:
	go mod tidy