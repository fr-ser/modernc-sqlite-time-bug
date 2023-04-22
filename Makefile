build:
	CGO_ENABLED=0 go build -o dist/main_local main.go 
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/main_darwin_amd64 main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o dist/main_darwin_arm64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/main_linux_amd64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o dist/main_linux_arm64 main.go
