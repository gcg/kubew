.PHONY: all ci build run update get test run lint tidy

run: 
	@go run -ldflags "-X main.version=local" github.com/gcg/kubew/cmd/kubew

tidy: 
	@go mod tidy

build: 
	@go build ./... 

get:
	@go get ./...  

update: 
	@go get -u ./... 	

test: 
	@go test ./... 

lint: 
	@go install honnef.co/go/tools/cmd/staticcheck@latest 
	@staticcheck ./... 
