mod:
	clear
	go mod tidy
	go mod vendor
	go fmt ./...

run.0:
	go run ./examples/screen/main.go
run.1:
	go run ./examples/free_pos_lit/main.go
run.2:
	go run ./examples/free_pos_string/main.go