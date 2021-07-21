gen:
	@go generate ./...

mod:
	@go mod tidy
	@go mod vendor
