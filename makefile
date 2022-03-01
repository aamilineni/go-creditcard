clean:
	go clean

test:
	go test -v

cover:
	go test -coverprofile cp.out
	go tool cover -html=cp.out