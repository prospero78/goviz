mod:
	clear
	go mod tidy
	go mod vendor
	go fmt ./...

run.screen:
	go run ./examples/screen/main.go