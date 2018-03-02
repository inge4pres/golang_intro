hello-go-local-build:
	go build ./cmd/hello_go -o hello_go

hello-go-docker:
	GOOS=linux GOARCH=amd64 go build -o app ./cmd/hello_go 
	docker build -t hello-go .
	docker run --rm hello-go
	file app
	rm -rf app
	
