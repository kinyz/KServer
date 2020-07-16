CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/socket socket/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/websocket websocket/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/chat chat/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/discovery discovery/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/lock lock/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/login login/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/oauth oauth/main.go
