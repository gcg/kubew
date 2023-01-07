.PHONY: all ci build run update get test run lint 

run: 
	@go run -ldflags "-X main.version=local" github.com/gcg/cmd/kubew 

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
