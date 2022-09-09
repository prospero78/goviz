mod:
	clear
	go mod tidy
	go mod vendor
	go fmt ./...