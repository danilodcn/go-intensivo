v_test:
	go clean -testcache
	go test -v ./...

test:
	go test ./...

tidy:
	go mod tidy